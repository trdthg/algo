create view aa as select ename, sal from emp;

create or replace trigger update_emp
after update on aa
for each row
begin
    insert into emp (ename, sal) values(ename, sal);
end;

create or replace function func(eename varchar)
return number is res number;
begin
    select sal into res from emp where ename = eename;
    return res;
end;

create or replace procedure procedure_test (
    noo out number
) is
begin
    select ename, sal from emp where empno = noo;
end;