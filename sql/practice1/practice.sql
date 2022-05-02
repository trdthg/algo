-- 1.查询" 01 "课程比" 02 "课程成绩高的学生的信息及课程分数

SELECT
	s.*,
	t1.SId,
	t1.Cname name1,
	t1.SCore SCore1,
	t2.Cname name2,
	t2.SCore SCore2
FROM
	`Student` s,
	(
		SELECT
			SC.SId,
			Course.Cname,
			SC.SCore
		FROM
			`SC`,
			Course
		WHERE
			SC.CId = '01'
			AND SC.CId = Course.CId
	) t1,
	(
		SELECT
			SC.SId,
			Course.Cname,
			SC.SCore
		FROM
			`SC`,
			Course
		WHERE
			SC.CId = '02'
			AND SC.CId = Course.CId
	) t2
WHERE
	t1.SId = t2.SId
	AND s.SId = t1.SId
	AND t1.SCore > t2.SCore;

-- 1.1 查询同时存在" 01 "课程和" 02 "课程的情况

SELECT
	s.*,
	t1.SId,
	t1.Cname name1,
	t1.SCore SCore1,
	t2.Cname name2,
	t2.SCore SCore2
FROM
	`Student` s,
	(
		SELECT
			SC.SId,
			Course.Cname,
			SC.SCore
		FROM
			`SC`,
			Course
		WHERE
			SC.CId = '02'
			AND SC.CId = Course.CId
	) t2
	LEFT JOIN (
		SELECT
			SC.SId,
			Course.Cname,
			SC.SCore
		FROM
			`SC`,
			Course
		WHERE
			SC.CId = '01'
			AND SC.CId = Course.CId
	) t1 -- WHERE
	ON t1.SId = t2.SId --   AND t1.SId = NULL
WHERE
	--   t1.SCore > t2.SCore
	s.SId = t1.SId;

--   AND t1.SCore < t2.SCore;

-- 1.2 查询存在" 01 "课程但可能不存在" 02 "课程的情况 (不存在时显示为 null )

SELECT
	s.*,
	t1.SId,
	t1.Cname name1,
	t1.SCore SCore1,
	t2.Cname name2,
	t2.SCore SCore2
FROM
	`Student` s,
	(
		SELECT
			SC.SId,
			Course.Cname,
			SC.SCore
		FROM
			`SC`,
			Course
		WHERE
			SC.CId = '01'
			AND SC.CId = Course.CId
	) t1
	LEFT JOIN (
		SELECT
			SC.SId,
			Course.Cname,
			SC.SCore
		FROM
			`SC`,
			Course
		WHERE
			SC.CId = '02'
			AND SC.CId = Course.CId
	) t2 ON t1.SId = t2.SId
WHERE
	s.SId = t1.SId;

-- 1.3 查询不存在" 01 "课程但存在" 02 "课程的情况

SELECT
	s.*,
	t2.Cname name2,
	t2.SCore SCore2
FROM
	`Student` s,
	(
		SELECT
			SC.SId,
			Course.Cname,
			SC.SCore
		FROM
			`SC`,
			Course
		WHERE
			SC.CId = '02'
			AND SC.CId = Course.CId
	) t2
WHERE
	s.SId NOT IN (
		SELECT
			SC.SId
		FROM
			SC
		WHERE
			SC.CId = '01'
	)
	AND s.SId = t2.SId;

-- 2. 查询平均成绩大于等于 60 分的同学的学生编号和学生姓名和平均成绩

SELECT
	*
FROM
	`Student` s,
	(
		SELECT
			SC.SId,
			COUNT(SC.CId) cnt,
			AVG(SC.SCore) avg
		FROM
			`SC`
		GROUP BY
			SC.SId --  统计课程数
		HAVING
			cnt = (
				SELECT
					COUNT(*)
				FROM
					`Course`
			) - 1
	) a
WHERE
	s.SId = a.SId
	AND a.avg > 60;

--   3. 查询在 SC 表存在成绩的学生信息

SELECT
	DISTINCT s.*
FROM
	`Student` s,
	`SC` SC
WHERE
	s.SId = SC.SId;

-- 4. 查询所有同学的学生编号、学生姓名、选课总数、所有课程的总成绩 (没成绩的显示为 null )

SELECT
	s.SId,
	s.Sname,
	SC.cnt,
	SC.sum
FROM
	`Student` s
	LEFT JOIN (
		SELECT
			SC.SId,
			COUNT(SC.CId) cnt,
			SUM(SC.SCore) sum
		FROM
			`SC`
		GROUP BY
			SC.SId
	) SC ON s.SId = SC.SId
WHERE
	s.SId = SC.SId;

-- 4.1 查有成绩的学生信息

SELECT
	s.SId,
	s.Sname,
	SC.cnt,
	SC.sum
FROM
	`Student` s
	RIGHT JOIN (
		SELECT
			SC.SId,
			COUNT(SC.CId) cnt,
			SUM(SC.SCore) sum
		FROM
			`SC`
		GROUP BY
			SC.SId
	) SC ON s.SId = SC.SId
WHERE
	s.SId = SC.SId;

-- 5. 查询「李」姓老师的数量

SELECT COUNT(TId) FROM `Teacher` WHERE Tname LIKE '李%';

-- 6. 查询学过「张三」老师授课的同学的信息

SELECT
	s.*
FROM
	`Student` s,
	`SC` SC
WHERE
	s.SId = SC.SId
	AND CId IN (
		SELECT
			CId
		FROM
			`Course`
		WHERE
			TId = (
				SELECT
					TId
				FROM
					`Teacher`
				WHERE
					Tname = '张三'
			)
	);

SELECT
	s.*
FROM
	`Student` s,
	`SC` SC
	RIGHT JOIN (
		SELECT
			CId
		FROM
			`Course`
		WHERE
			TId = (
				SELECT
					TId
				FROM
					`Teacher`
				WHERE
					Tname = '张三'
			)
	) a ON SC.CId = a.CId
WHERE
	s.SId = SC.SId;

-- 7. 查询没有学全所有课程的同学的信息

SELECT SC.SId, COUNT(SC.CId) cnt FROM `SC` SC GROUP BY SId;

SELECT
	s.*,
	SC.cnt
FROM
	`Student` s
	LEFT JOIN (
		SELECT
			SC.SId,
			COUNT(SC.CId) cnt
		FROM
			`SC` SC
		GROUP BY
			SId
	) SC ON SC.SId = s.SId
WHERE
	SC.cnt IS NULL
	OR SC.cnt < 3;

-- WHERE

--   s.SId = SC.SId;

--   SC.cnt != (

--     SELECT

--       COUNT(*)

--     FROM

--       `Course`

--   );

SELECT
	--   DISTINCT s.*
	s.SId,
	COUNT(s.SId) AS `缺少的门数`
FROM
	`Student` s,
	(
		SELECT
			t1.SId
		FROM
			(
				SELECT
					*
				FROM
					`Student` s,
					`Course` c
			) AS t1
			LEFT JOIN `SC` t2 ON t1.SId = t2.SId
			AND t1.CId = t2.CId
		WHERE
			t2.SId IS NULL
	) SC
WHERE
	s.SId = SC.SId
GROUP BY
	SId;

-- 8. 查询至少有一门课与学号为" 01 "的同学所学相同的同学的信息

SELECT
	DISTINCT SC2.SId,
	s.Sname
FROM
	`Student` s,
	`SC` SC1,
	`SC` SC2
WHERE
	SC1.SId = 01
	AND SC1.CId = SC2.CId -- 同一个课
	AND SC1.SId != SC2.SId -- 不是同一个人
	AND s.SId = SC2.SId;

-- 它的信息

SELECT
	DISTINCT s.SId,
	s.Sname
FROM
	`Student` s,
	`SC` SC
WHERE
	SC.CId IN (
		SELECT
			SC.CId
		FROM
			`SC`
		WHERE
			SC.SId = '01'
	)
	AND s.SId != '01'
	AND s.SId = SC.SId;

-- 9. 查询和" 01 "号的同学学习的课程 完全相同的其他同学的信息

-- 9.1 查询有缺课的（NULL）

SELECT
	*
FROM
	`Student`
WHERE
	SId NOT IN (
		SELECT
			DISTINCT t1.SId
		FROM
			(
				SELECT
					s.SId,
					sc.CId
				FROM
					`Student` s
					INNER JOIN (
						SELECT
							sc.CId
						FROM
							`SC` sc
						WHERE
							sc.SId = '01'
					) sc
			) t1
			LEFT JOIN `SC` t2 ON t1.SId = t2.SId -- 魔幻
			AND t1.CId = t2.CId
		WHERE
			t2.SId IS NULL
	);

-- 10. 查询没学过"张三"老师讲授的任一门课程的学生姓名

-- 查询 NOT IN 学过张三老师的课 的同学

-- 张三老师教的课

SELECT
	CId
FROM
	`Course`
WHERE
	TId = (
		SELECT
			TId
		FROM
			`Teacher`
		WHERE
			Tname = '张三'
	);

SELECT
	*
FROM
	`Student` s
	LEFT JOIN (
		SELECT
			sc.SId,
			sc.CId
		FROM
			`SC` sc
	) sc ON s.SId = sc.SId
WHERE
	sc.CId IN (
		SELECT
			CId
		FROM
			`Course`
		WHERE
			TId = (
				SELECT
					TId
				FROM
					`Teacher`
				WHERE
					Tname = '张三'
			)
	);

-- 11. 查询两门及其以上不及格课程的同学的学号，姓名及其平均成绩

SELECT
	s.SId,
	s.Sname,
	sc.avg
FROM
	`Student` s,
	(
		SELECT
			sc.SId,
			AVG(sc.score) avg
		FROM
			`SC` sc,
			(
				-- 查询不及格>=两次的学生的学号
				SELECT
					sc.SId,
					COUNT(*) cnt
				FROM
					`SC` sc
				WHERE
					sc.score < 60
				GROUP BY
					sc.SId
			) sc2
		WHERE
			sc2.cnt >= 2
			AND sc.SId = sc2.SId
		GROUP BY
			sc.SId
	) sc
WHERE
	s.SId = sc.SId;

SELECT
	s.SId,
	--   s.Sname,
	AVG(sc.score) avg
FROM
	`Student` s,
	`SC` sc
WHERE
	s.SId = sc.SId
	AND sc.score < 60
GROUP BY
	sc.SId
HAVING
	COUNT(*) >= 2;

-- 12. 检索" 01 "课程分数小于 60，按分数降序排列的学生信息

SELECT
	*
FROM
	`Student` s,
	`SC` sc
WHERE
	sc.CId = '01'
	AND sc.score < 60
	AND s.SId = sc.SId
ORDER BY
	sc.score;

-- 13. 按平均成绩从高到低显示所有学生的所有课程的成绩以及平均成绩

SELECT
	s.SId,
	sc.Sname,
	sc.Cname,
	sc.score,
	-- s.Sname,
	s.avg
FROM
	(
		SELECT
			s.Sname,
			sc.*,
			c.Cname
		FROM
			`Student` s,
			`SC` sc,
			`Course` c
		WHERE
			sc.CId = c.CId
			AND s.SId = sc.SId
	) sc
	LEFT JOIN (
		SELECT
			sc.SId,
			-- sc.CId,
			AVG(sc.score) as avg
		FROM
			`SC` sc
		GROUP BY
			sc.SId
	) s ON s.SId = sc.SId
ORDER BY
	sc.score DESC;

-- 14. 查询各科成绩最高分、最低分和平均分：

SELECT
	sc.CId,
	MAX(sc.score),
	MIN(sc.score),
	AVG(sc.score)
FROM
	`SC` sc
GROUP BY
	sc.CId;

-- 以如下形式显示：课程 ID，课程 name，最高分，最低分，平均分，及格率，中等率，优良率，优秀率，及格为>=60，中等为：70-80，优良为：80-90，优秀为：>=90

-- 要求输出课程号和选修人数，查询结果按人数降序排列，若人数相同，按课程号升序排列

SELECT
	c.CId,
	c.Cname,
	sc.max,
	sc.min,
	sc.avg,
	sc.cnt,
	sc.a,
	sc.b,
	sc.c,
	sc.d
FROM
	(
		SELECT
			sc.CId,
			MAX(sc.score) max,
			MIN(sc.score) min,
			AVG(sc.score) avg,
			COUNT(sc.SId) cnt,
			SUM(
				CASE
					when sc.score >= 60 then 1
					else 0
				END
			) / COUNT(*) a,
			SUM(
				CASE
					when sc.score >= 70
					and sc.score < 80 then 1
					else 0
				END
			) / COUNT(*) b,
			SUM(
				CASE
					when sc.score >= 80
					and sc.score < 90 then 1
					else 0
				END
			) / COUNT(*) c,
			SUM(
				CASE
					when sc.score >= 90 then 1
					else 0
				END
			) / COUNT(*) d
		FROM
			`SC` sc
		GROUP BY
			sc.CId
	) sc,
	`Course` c
WHERE
	sc.CId = c.CId
ORDER BY
	sc.cnt DESC,
	sc.CId ASC;

-- 15. 按各科成绩进行排序，并显示排名， SCore 重复时保留名次空缺

SELECT
	sc.SId,
	sc.CId,
	sc.score,
	t.rank + 1 rank
FROM
	`SC` sc
	LEFT JOIN (
		-- 依次用每一个成绩，然后用这个成绩取查询比他分高的有几个，就是他的排名
		SELECT
			sc.SId,
			sc.CId,
			(
				SELECT
					COUNT(sc2.score)
				FROM
					`SC` sc2
				WHERE
					sc.CId = sc2.CId
					AND sc.score < sc2.score
			) rank
		FROM
			`SC` sc
	) t ON sc.SId = t.SId
	AND sc.CId = t.CId
ORDER BY
	sc.CId,
	t.rank;

-- 15.1 按各科成绩进行排序，并显示排名， SCore 重复时合并名次

SELECT
	sc.SId,
	sc.CId,
	sc.score,
	t.rank + 1 rank
FROM
	`SC` sc
	LEFT JOIN (
		SELECT
			sc.SId,
			sc.CId,
			(
				SELECT
					-- 这里去重
					COUNT(DISTINCT sc2.score)
				FROM
					`SC` sc2
				WHERE
					sc.CId = sc2.CId
					AND sc.score < sc2.score
			) rank
		FROM
			`SC` sc
	) t ON sc.SId = t.SId
	AND sc.CId = t.CId
ORDER BY
	sc.CId,
	t.rank;

-- 16. 查询学生的总成绩，并进行排名，总分重复时保留名次空缺

-- SELECT

-- 	sc.SId,

-- 	-- sc.CId,

-- 	-- sc.score,

-- 	t.rank + 1 rank

-- FROM

-- 	`SC` sc

-- 	LEFT JOIN (

SELECT
	tt.SId,
	tt.sum,
	(
		SELECT
			-- 这里去重
			COUNT(DISTINCT t2.sum)
		FROM
			(
				SELECT
					sc.SId,
					SUM(sc.score) sum
				FROM
					`SC` sc
				GROUP BY
					sc.SId
			) t2
		WHERE
			tt.sum < t2.sum
	) + 1 rank
FROM
	(
		SELECT
			sc.SId,
			SUM(sc.score) sum
		FROM
			`SC` sc
		GROUP BY
			sc.SId
	) tt
ORDER BY
	rank;

select
	k1.*,
	count(k2.total) rank
from
	(
		select
			SId,
			sum(score) total
		from
			`SC` sc
		group by
			SId
	) k1,
	(
		select
			sum(score) total
		from
			SC sc
		group by
			SId
	) k2
where
	k2.total >= k1.total
group by
	k1.SId
order by
	rank;

-- 16.1 查询学生的总成绩，并进行排名，总分重复时不保留名次空缺

-- 17. 统计各科成绩各分数段人数：课程编号，课程名称，[100-85]，[85-70]，[70-60]，[60-0] 及所占百分比

-- 18. 查询各科成绩前三名的记录

SELECT
	sc.SId,
	sc.CId,
	sc.score
FROM
	`SC` sc
WHERE
	(
		SELECT
			COUNT(*)
		FROM
			`SC` sc2
		WHERE
			sc.CId = sc2.CId
			AND sc.score < sc2.score
	) < 3
ORDER BY
	sc.CId,
	sc.score DESC;

-- 19. 查询每门课程被选修的学生数

SELECT sc.CId, COUNT(sc.SId) FROM `SC` sc GROUP BY sc.CId ;

-- 20. 查询出只选修两门课程的学生学号和姓名

SELECT
	*
FROM(
		SELECT
			sc.SId,
			COUNT(sc.CId) cnt
		FROM
			`SC` sc
		GROUP BY
			sc.SId
	) t
WHERE
	t.cnt = 2;

-- 21. 查询男生、女生人数

SELECT s.Ssex, COUNT(s.SId) FROM `Student` s GROUP BY s.Ssex;

-- 22. 查询名字中含有「风」字的学生信息

SELECT * FROM `Student` s WHERE s.Sname LIKE '%风%' ;

-- 23. 查询同名同性学生名单，并统计同名人数

SELECT
	*
FROM
	`Student` s,
	(
		SELECT
			s.Sname,
			s.Ssex,
			COUNT(*) cnt
		FROM
			`Student` s
		GROUP BY
			s.Sname,
			s.Ssex
		HAVING
			cnt > 1
	) t
WHERE
	s.Sname = t.Sname
	AND s.Ssex = t.Ssex;

-- 24. 查询 1990 年出生的学生名单

SELECT * FROM `Student` s WHERE s.Sage LIKE '1990%';

-- 25. 查询每门课程的平均成绩，结果按平均成绩降序排列，平均成绩相同时，按课程编号升序排列

SELECT
	sc.CId,
	AVG(sc.score) avg
FROM
	`SC` sc
GROUP BY
	sc.CId
ORDER BY
	avg DESC,
	sc.CId;

-- 26. 查询平均成绩大于等于 85 的所有学生的学号、姓名和平均成绩

SELECT
	*
FROM
	`Student` s,
	(
		SELECT
			sc.SId,
			AVG(sc.score) avg
		FROM
			`SC` sc
		GROUP BY
			sc.SId
		HAVING
			avg >= 85
	) sc
WHERE
	s.SId = sc.SId;

-- 27. 查询课程名称为「数学」，且分数低于 60 的学生姓名和分数

SELECT
	*
FROM
	`Student` s,
	`SC` sc
WHERE
	sc.CId = (
		SELECT
			c.CId
		FROM
			`Course` c
		WHERE
			c.Cname = '数学'
	)
	AND sc.score < 60
	AND s.SId = sc.SId;

-- 28. 查询所有学生的课程及分数情况（存在学生没成绩，没选课的情况）

SELECT
	s.SId,
	s.CId,
	sc.score
FROM
	(
		SELECT
			s.SId,
			c.CId
		FROM
			(`Student` s, `Course` c)
	) s
	LEFT JOIN `SC` sc ON s.SId = sc.SId;

SELECT
	s.SId,
	max(
		case
			when sc.CId = '01' then sc.score
			else null
		end
	) as '语文'
FROM
	`Student` s
	LEFT JOIN `SC` sc ON s.SId = sc.SId
GROUP BY
	s.SId;

-- 29. 查询任何一门课程成绩在 70 分以上的姓名、课程名称和分数

select
	s.SId,
	s.Sname,
	c.Cname,
	sc.score
from
	Student s,
	SC sc,
	Course c,
	(
		select
			SId
		from
			SC sc
		group by
			SId
		having
			min(score) > 70
	) n
where
	n.SId = s.SId
	and sc.CId = c.CId
	and n.SId = sc.SId;

-- 30. 查询不及格的课程

SELECT DISTINCT sc.SId FROM `SC` sc WHERE sc.score < 60;

-- 31. 查询课程编号为 01 且课程成绩在 80 分以上的学生的学号和姓名

SELECT
	DISTINCT sc.SId,
	s.Sname,
	sc.CId,
	sc.score
FROM
	`Student` s,
	`SC` sc
WHERE
	s.SId = sc.SId
	AND sc.CId = '01'
	AND sc.score >= 60;

-- 32. 求每门课程的学生人数

SELECT CId, COUNT(*) FROM `SC` GROUP BY CId;

-- 33. 成绩不重复，查询选修「张三」老师所授课程的学生中，成绩最高的学生信息及其成绩

SELECT
	*
FROM
	`Student` s,(
		SELECT
			*
		FROM
			`SC` sc
		WHERE
			sc.CId = (
				SELECT
					CId
				FROM
					`Course` c
				WHERE
					c.TId = (
						SELECT
							TId
						FROM
							`Teacher`
						WHERE
							Tname = '张三'
					)
			)
		ORDER BY
			sc.score DESC
		LIMIT
			1
	) sc
WHERE
	s.SId = sc.SId;

-- 34. 成绩有重复的情况下，查询选修「张三」老师所授课程的学生中，成绩最高的学生信息及其成绩

SELECT
	*
FROM
	`SC` sc,
	(
		SELECT
			sc.CId,
			MAX(sc.score) max
		FROM
			`SC` sc
		WHERE
			sc.CId = (
				SELECT
					CId
				FROM
					`Course` c
				WHERE
					c.TId = (
						SELECT
							TId
						FROM
							`Teacher`
						WHERE
							Tname = '张三'
					)
			)
		GROUP BY
			sc.CId
	) t
WHERE
	sc.CId = t.CId
	AND sc.score = t.max;

SELECT
	sc.CId
FROM
	`SC` sc,
	MAX(sc.score) max
WHERE
	sc.CId IN (
		SELECT
			CId
		FROM
			`Course` c
		WHERE
			c.TId = (
				SELECT
					TId
				FROM
					`Teacher`
				WHERE
					Tname = '张三'
			)
	)
	AND sc.score = max -- AND (
	-- 	SELECT
	-- 		*
	-- 	FROM
	-- 		`SC` sc2
	-- 	WHERE
	-- 		sc.CId = sc2.CId
	-- 		AND sc.score < sc2.score
	-- ) < 2
	-- ORDER BY
	-- 	sc.score DESC
GROUP BY
	sc.CId;

-- 35. 查询不同课程成绩相同的学生的学生编号、课程编号、学生成绩

SELECT
	DISTINCT sc1.*
FROM
	`SC` sc1,
	`SC` sc2
WHERE
	sc1.SId = sc2.SId
	AND sc1.score = sc2.score
	AND sc1.CId != sc2.CId;

-- 36. 查询每门功成绩最好的前两名

SELECT
	*
FROM
	`SC` sc
WHERE
	(
		SELECT
			COUNT(*)
		FROM
			`SC` sc2
		WHERE
			sc.CId = sc2.CId
			AND sc.score < sc2.score
	) < 2
ORDER BY
	sc.CId,
	sc.score;

-- 37. 统计每门课程的学生选修人数（超过 5 人的课程才统计）。

SELECT
	sc.CId,
	COUNT(*)
FROM
	`SC` sc
GROUP BY
	sc.CId
HAVING
	COUNT(*) > 5;

-- 38. 检索至少选修两门课程的学生学号

SELECT sc.SId FROM `SC` sc GROUP BY sc.SId HAVING COUNT(*) >= 3;

-- 39. 查询选修了全部课程的学生信息

SELECT
	s.*
FROM
	`Student` s,(
		SELECT
			sc.SId
		FROM
			`SC` sc
		GROUP BY
			sc.SId
		HAVING
			COUNT(*) = (
				SELECT
					COUNT(*)
				FROM
					`Course`
			)
	) t
WHERE
	s.SId = t.SId;

-- 40. 查询各学生的年龄，只按年份来算

SELECT
	s.SId,
	s.Sname,
	TIMESTAMPDIFF(YEAR, s.Sage, NOW()) age
FROM
	`Student` s;

-- 41. 按照出生日期来算，当前月日 < 出生年月的月日则，年龄减一

SELECT
	s.SId,
	s.Sname,
	TIMESTAMPDIFF(DAY, s.Sage, NOW()) age
FROM
	`Student` s;

-- 42. 查询本周过生日的学生

-- 43. 查询下周过生日的学生

-- 44. 查询本月过生日的学生

-- 45. 查询下月过生日的学生