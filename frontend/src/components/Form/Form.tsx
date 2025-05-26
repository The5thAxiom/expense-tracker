import { useEffect, useReducer, useState } from "react";
import { ExpenseFormData, formReducer, initialFormState } from "./FormReducer";

type FormProps = {
    onSubmit: (p: ExpenseFormData) => void;
};

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

function Form({ onSubmit }: FormProps) {
    const [currencies, setCurrencies] = useState<Currency[] | null>(
        [{id: "INR", name: "Indian Rupee", symbol: "₹"}]
    );
    const [categories, setCategories] = useState<Category[] | null>(null);
    const [subCategories, setSubCategories] = useState<SubCategory[] | null>(null);
    const [purposes, setPurposes] = useState<Purpose[] | null>(null);

    const [errorMessage, setErrorMessage] = useState<string | null>(null);

    async function fetchCategories() {
        const res = await fetch('/api/categories');
        const json = await res.json();
        if (json.errorMessage) {
            setErrorMessage(json.errorMessage);
            return;
        }

        setCategories(json.data);
    }

    async function fetchCurrencies() {
        const res = await fetch('/api/currencies');
        const json = await res.json();
        if (json.errorMessage) {
            setErrorMessage(json.errorMessage);
            return;
        }

        setCurrencies(json.data);
    }

    async function fetchSubCategoreis(categoryId: string) {
        const res = await fetch(`/api/categories/${categoryId}/sub-categories`);
        const json = await res.json();
        if (json.errorMessage) {
            setErrorMessage(json.errorMessage);
            return;
        }

        setSubCategories(json.data);
    }

    async function fetchPurposes() {
        const res = await fetch('/api/purposes');
        const json = await res.json();
        if (json.errorMessage) {
            setErrorMessage(json.errorMessage);
            return;
        }

        setPurposes(json.data);
    }

    // useEffect(() => {
    //     fetchCategories();
    //     fetchCurrencies();
    //     fetchPurposes();
    // }, [])

    const [formData, formDataDispatch] = useReducer(formReducer, initialFormState);

    // useEffect(() => {
    //     if (payment.categoryId && payment.categoryId.length > 0) {
    //         fetchSubCategoreis(payment.categoryId)
    //     }
    // }, [payment.categoryId])

    useEffect(() => {
        console.log(formData)
    }, [formData])

    function handleInputChange(e: React.ChangeEvent<HTMLInputElement>) {
        formDataDispatch({type: "CHANGE_INPUT", payload: e})
    }

    function handleSelectChange(e: React.ChangeEvent<HTMLSelectElement>) {
        formDataDispatch({type: "CHANGE_SELECT", payload: e})
    }

    return <>
        <form onSubmit={e => {
            e.preventDefault();
            onSubmit(formData);
        }}>
            <div>{errorMessage}</div>
            <div>
                <label htmlFor='date'>Expense Date:</label>
                <input name='date' id='date-input' type='date' onChange={handleInputChange} value={formData.date.toISOString().substring(0, 10)} />
            </div>
            <div>
                <label htmlFor='amount'>Expense Amount:</label>
                <input name='amount' type='number' onChange={handleInputChange} value={formData.amount} />
            </div>
            <div>
                <label htmlFor='currency'>Select a Currency:</label>
                {currencies && <>
                    <select name='currencyId' defaultValue='' onChange={handleSelectChange}>
                        <option value=''>New Currency</option>
                        {currencies.map(c => <option key={c.id} value={c.id}>
                            {c.name} ({c.id} | {c.symbol})
                        </option>)}
                    </select>
                </>
                }
            </div>
            <div>
                <label htmlFor='description'>Description:</label>
                <input name='description' type='text' onChange={handleInputChange} value={formData.description} />
            </div>
            <div>
                <label htmlFor='category'>Select a Category:</label>
                {/* {categories && <select name='category' defaultValue='' onChange={handleSelectChange}>
                    <option value=''>New Category</option>
                    {categories.map(c => <option key={c.id} value={c.id}>
                        {c.name}{c.description && <> ({c.description})</>}
                    </option>)}
                </select>} */}
                <input name='newCategoryName' type='text' onChange={handleInputChange} value={formData.newCategoryName} />
            </div>
            {(formData.categoryId || formData.newCategoryName) && <div>
                <label htmlFor='subCategory'>Select a sub category:</label>
                {/* {subCategories && <select name='sub category' defaultValue='' onChange={handleSelectChange}>
                    <option value=''>New Sub Category</option>
                    {subCategories.map(sc => <option key={sc.id} value={sc.id}>
                        {sc.name}{sc.description && <> ({sc.description})</>}
                    </option>)}
                </select>} */}
                <input name='newSubCategoryName' type='text' onChange={handleInputChange} value={formData.newSubCategoryName} />
            </div>}
            <div>
                <label htmlFor='purpose'>Select a Purpose:</label>
                {/* {purposes && <select name='purpose' defaultValue='' onChange={handleSelectChange}>
                    <option value=''>New purpose</option>
                    {purposes.map(c => <option key={c.id} value={c.id}>
                        {c.name}{c.description && <> ({c.description})</>}
                    </option>)}
                </select>} */}
            </div>
            <div>
                <label htmlFor='notes'>Notes:</label>
                <input name='notes' type='text' onChange={handleInputChange} value={formData.notes} />
            </div>
            <button type='submit'>Submit</button>
        </form>
    </>;
}

export default Form;
