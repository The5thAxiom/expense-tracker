import { useEffect, useState } from "react";
import Form from "../../components/Form/Form";

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
  id: string;
  name?: string;
  symbol?: string;
};

type Expense = {
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


const Expenses = () => {
  const [payments, setExpenses] = useState<Expense[]>([]);
  const [errorMessage, setErrorMessage] = useState<string | null>(null);

  async function fetchPayements() {
    const res = await fetch("/api/expenses", {
      headers: {
        'Authorization': `Bearer ${localStorage.getItem('token')}`,
        'Content-Type': 'application/json'
      },
      method: 'GET',
    });
    const json = await res.json();

    if (json.errorMessage) {
      setErrorMessage(json.errorMessage);
      return;
    }

    setExpenses(json.data);
  }

  useEffect(() => {
    fetchPayements();
  }, []);


  return (
    <main>
      <Form onSubmit={async e => {

        const efinal = {
          date: e.date.toISOString().substring(0, 10),
          description: e.description,
          amount: e.amount * 100, // cause we are storing amount in the lowest denomination
          notes: !e.notes ? null : e.notes,
          currencyId: "INR",
          category: {
            isNew: true,
            name: e.newCategoryName,
          },
          subCategory: {
            isNew: true,
            name: e.newSubCategoryName
          },
          // purpose: {
          //     isNew: true,
          //     name: e.newPurposeName ?? 'abc',
          //     description: "Personal expenses"
          // },
          tags: [
            {
              isNew: true,
              name: "Urgent",
              description: "Urgent expenses"
            }
          ]
        };

        const res = await fetch("/api/expenses", {
          headers: {
            'Authorization': `Bearer ${localStorage.getItem('token')}`,
            'Content-Type': 'application/json'
          },
          method: 'POST',
          body: JSON.stringify(efinal),
        });
        const data = await res.json();
        console.log(data);
      }} />
      <br />
      <button onClick={fetchPayements}> Fetch </button>
      {errorMessage}
      <table>
        <thead>
          <tr>
            {/* <th>Id</th> */}
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
              {/* <td>{payment.id}</td> */}
              <td>{new Date(payment.date).toDateString()}</td>
              <td>
                {new Intl.NumberFormat('en-IN', {
                  style: 'currency',
                  currency: payment.currency.id
                }).format(payment.amount / 100)}
              </td>
              <td>{payment.description}</td>
              <td>{payment.category.name}</td>
              <td>{payment.subCategory.name}</td>
              <td>{payment.purpose?.name ?? '-'}</td>
              <td>{payment.notes ?? '-'}</td>
            </tr>;
          }).reverse()}
        </tbody>
      </table>
    </main>
  )

}

export default Expenses;