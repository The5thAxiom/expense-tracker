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
    type: "CHANGE_INPUT", payload: React.ChangeEvent<HTMLInputElement>
} | {
    type: "CHANGE_SELECT", payload: React.ChangeEvent<HTMLSelectElement>
}

export function formReducer(state: ExpenseFormData, { type, payload }: ExpenseFormAction): ExpenseFormData {
    const name = payload.target.name;

    switch (type) {
        case "CHANGE_INPUT":
            const inputType = payload.target.type;
            let inputValue;

            if (inputType === "number") {
                inputValue = payload.target.valueAsNumber
            } else if (inputType === "date") {
                inputValue = payload.target.valueAsDate
            } else {
                inputValue = payload.target.value
            }

            switch (inputType) {
                case "number":
                    inputValue = payload.target.valueAsNumber
                    break;
                case "date":
                    inputValue = payload.target.valueAsDate
                    break;
                default:
                    inputValue = payload.target.value
            }

            return { ...state, [name]: inputValue }
        case 'CHANGE_SELECT':
            return { ...state, [name]: payload.target.value }
        default:
            return state
    }
}