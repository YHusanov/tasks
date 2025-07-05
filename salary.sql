SELECT salary
FROM (SELECT salary, ROW_NUMBER() OVER (ORDER BY salary DESC) as n
         FROM (SELECT DISTINCT salary FROM Employee)
     )
WHERE n = 1;