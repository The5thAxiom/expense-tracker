import Form from "../../components/Form/Form";

const Home = () => {
    return (
        <main>
            <h1>Home</h1>
            <Form onSubmit={async e => {

                const efinal = {
                    date: new Date().toISOString(),
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
        </main>
    )
}

export default Home;