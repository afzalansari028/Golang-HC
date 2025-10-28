## IMP SQL Queries

```sql
### 1️⃣ Highest Salary from Each Department  
This query retrieves the highest salary for each department by joining the `employee` and `department` tables.

select d.department_name, MAX(e.employee_salary)
from employee e
join department d on e.department_id = d.department_id
GROUP BY d.department_name;
---------------------------------
## Included employee name column
select d.department_name, e.employee_name, e.employee_salary
from employee e
join department d on e.department_id = d.department_id
where e.employee_salary = (
    select MAX(e1.employee_salary)
    from employee e1
    where e1.department_id = e.department_id
);
