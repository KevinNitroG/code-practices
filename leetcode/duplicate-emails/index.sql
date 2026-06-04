-- https://leetcode.com/problems/duplicate-emails

SELECT
  email AS Email
FROM
  Person
GROUP BY
  email
HAVING
  COUNT(email) > 1;
