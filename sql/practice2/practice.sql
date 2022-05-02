SELECT
	name,
	score,
	CASE
		WHEN @preScore = score THEN @curRank
		WHEN @preScore: = score THEN @curRank: = @curRank + 1
	END AS `Rank`
FROM
	test1,(
		SELECT
			@curRank: = 0 AS curRank,
			@preScore: = NULL AS preRank
	) i
WHERE
	sex = 1
ORDER BY
	score DESC;