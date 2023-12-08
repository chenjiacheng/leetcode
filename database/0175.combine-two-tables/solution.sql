-- 解题思路：
-- 使用 left join

select FirstName, LastName, City, State
from Person left join Address
on Person.PersonId = Address.PersonId