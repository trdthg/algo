-- select

-- 	t1.university,

-- 	t2.difficult_level,

-- 	t2.avg_answer_cnt avg_answer_cnt

-- from

-- 	(

-- 		select

-- 			u.university,

-- 			count(*) cnt

-- 		from

-- 			user_profile u

-- 		group by

-- 			u.university

-- 	) t1,

-- 	(

select
	u.university,
	q.difficult_level,
	count(q.question_id) / count(distince u.device_id) avg_answer_cnt
from
	(
		select
			q.device_id,
			q.question_id,
			q2.difficult_level
		from
			question_practice_detail q,
			question_detail q2
		where
			q.question_id = q2.question_id
	) q,
	user_profile u
where
	u.device_id = q.device_id
group by
	u.university,
	q.difficult_level;

-- 	) t2

-- where

-- 	t1.university = t2.university;