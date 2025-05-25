import React from 'react';

export type ExpenseFormData = {
    date: Date;
    description: string;
    amount: number;
    currencyId: string | undefined;
    newCurrencyId: string | undefined;
    categoryId: string | undefined;
    newCategoryName: string | undefined;
    subCategoryId: string | undefined;
    newSubCategoryName: string | undefined;
    purposeId: string | undefined;
    newPurposeName: string | undefined;
    notes: string
};

export const initialFormState: ExpenseFormData = {
    date: new Date(),
    description: "",
    amount: 0,
    currencyId: 'INR',
    newCurrencyId: "",
    categoryId: "",
    newCategoryName: "",
    subCategoryId: "",
    newSubCategoryName: "",
    purposeId: "",
    newPurposeName: "",
    notes: "",
}

export type ExpenseFormAction = {
    type: "CHANGE_TEXT_INPUT", payload: React.ChangeEvent<HTMLInputElement>
} | {
    type: "CHANGE_DATE_INPUT", payload: React.ChangeEvent<HTMLInputElement>
}

export function formReducer(state: ExpenseFormData, action: ExpenseFormAction): ExpenseFormData {
    const type = action.payload.target.type;
    const name = action.payload.target.name;

    switch (action.type) {
        case "CHANGE_TEXT_INPUT":
            console.log(action.payload.target)
            return {
                ...state,
                [name]: type === 'number'
                    ? action.payload.target.valueAsNumber
                    : type === 'date'
                        ? action.payload.target.valueAsDate
                        : action.payload.target.value
            }
        case 'CHANGE_DATE_INPUT':
        default:
            return state
    }
}