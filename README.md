# Expense Tracker

## Running Import From Excel

```bash
cd backend
```

```bash
go run . import-excel --db ../database.db --reset-db --excel ../BudgetAndExpenses.xlsx --sheet Expenses
```

## Running Backend Server

```bash
cd backend
```

```bash
go run . serve --db ../database.db
```
