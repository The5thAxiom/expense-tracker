import { useEffect, useState } from 'react';
import './App.css'

type Category = {
  id: string;
  name: string;
  description?: string;
}

type SubCategory = {
  id: string;
  name: string;
  description?: string;
}

type Purpose = {
  id: string;
  name: string;
  description?: string;
}

type Currency = {
  abbreviation: string;
  name?: string;
  symbol?: string;
};

type Payment = {
  id: number;
  date: string;
  paymentIndex: number;
  description: string;
  amount: number;
  currency: Currency;
  category: Category;
  subCategory: SubCategory;
  purpose?: Purpose;
  notes?: string;
};


function App() {
  const [payments, setPayments] = useState<Payment[] | null>(null);
  const [errorMessage, setErrorMessage] = useState<string | null>(null);

  async function fetchPayements() {
    const res = await fetch("/api/payments");
    const json: { errorMessage: string, data: Payment[] } = await res.json();

    if (json.errorMessage) {
      setErrorMessage(json.errorMessage);
      return;
    }

    setPayments(json.data.toReversed());
  }

  useEffect(() => {
    fetchPayements();
  }, []);


  return (
    <>
      {errorMessage && <div>An error occured: {errorMessage}</div>}
      {!payments
        ? <div>Loading...</div>
        : payments.length === 0
          ? <div>No payments found :(, click the + button to add a payment</div>
          : <table>
            <thead>
              <tr>
                <th>Id</th>
                <th>Date</th>
                <th>Amount</th>
                <th>Description</th>
                <th>Category</th>
                <th>Subcategory</th>
                <th>Purpose</th>
                <th>Notes</th>
              </tr>
            </thead>
            <tbody>
              {payments.map((payment) => {
                return <tr key={payment.id}>
                  <td>{payment.id}</td>
                  <td>{new Date(payment.date).toDateString()}</td>
                  <td>{payment.currency.symbol} {payment.amount}</td>
                  <td>{payment.description}</td>
                  <td>{payment.category.name}</td>
                  <td>{payment.subCategory.name}</td>
                  <td>{payment.purpose?.name ?? '-'}</td>
                  <td>{payment.notes ?? '-'}</td>
                </tr>;
              })}
            </tbody>
          </table>
      }
    </>
  )
}

export default App
