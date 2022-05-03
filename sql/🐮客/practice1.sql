-- 25 UNION

select
	device_id,
	gender,
	age,
	gpa
from
	user_profile
where
	university = '山东大学'
union all
select
	device_id,
	gender,
	age,
	gpa
from
	user_profile
where
	gender = 'male';

-- 27 条件函数

select
	device_id,
	gender,
	(
		case
			when age < 20 then '20 岁以下'
			when age < 25 then '20-24 岁'
			when age >= 25 then '25 岁及以上'
			else '其他'
		end
	) age_cut
from
	user_profile;

SELECT
	IF(
		age < 25
		OR age IS NULL,
		'25 岁以下',
		'25 岁以及上'
	) age_cut,
	COUNT(device_id) Number
FROM
	user_profile
GROUP BY
	age_cut;

-- 28 日起函数

select
	day(date) day,
	count(*) question_cnt
from
	question_practice_detail
where
	year(date) = 2021
	and month(date) = 8
group by
	day -- 29 计算用户的平均次日留存率
	-- 题解 1
select
	count(date2) / count(date1) as avg_ret
from
	(
		select
			distinct qpd.device_id,
			qpd.date as date1,
			uniq_id_date.date as date2
		from
			question_practice_detail as qpd
			left join(
				select
					distinct device_id,
					date
				from
					question_practice_detail
			) as uniq_id_date on qpd.device_id = uniq_id_date.device_id
			and date_add(qpd.date, interval 1 day) = uniq_id_date.date
	) as id_last_next_date -- 题解 2
select
	avg(if(datediff(q.date2, q.date1) = 1, 1, 0)) as avg
from
	(
		select
			q.device_id,
			q.date date1,
			(
				lead(date, 1, null) over (
					partition by device_id
					order by
						date
				)
			) date2
		from
			(
				select
					distinct q.device_id,
					date
				from
					question_practice_detail q
			) q
	) q -- 30.字符串切割
select
	substring_index(u.profile, ',', -1) gender,
	count(*) number
from
	user_submit u
group by
	gender