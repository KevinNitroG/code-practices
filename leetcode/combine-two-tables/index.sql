-- https://leetcode.com/problems/combine-two-tables

SELECT
  p.firstName,
  p.lastName,
  s.city,
  s.state
FROM
  Person p
LEFT JOIN
  Address s
ON
  p.personId = s.personId;
