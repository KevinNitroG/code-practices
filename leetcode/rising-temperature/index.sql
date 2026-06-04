-- https://leetcode.com/problems/rising-temperature/

SELECT
  n.id
FROM
  Weather AS n
CROSS JOIN
  Weather AS p
WHERE
  n.recorddate - p.recorddate = 1
  AND n.temperature > p.temperature;
