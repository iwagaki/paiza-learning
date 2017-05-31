SELECT division_name, division_count
FROM (
SELECT division.division_id, division.division_name, count(*) AS division_count
FROM member JOIN division
ON member.division_id = division.division_id
GROUP BY division.division_name, division.division_id
ORDER BY member.division_id ASC
) AS t
