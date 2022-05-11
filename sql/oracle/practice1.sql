-- 1. 建表
DROP TABLE sc;

DROP TABLE student;

DROP TABLE course;

CREATE Table student(
    sno NUMBER(5) PRIMARY KEY,
    sname VARCHAR2(20),
    sex CHAR(3) NOT NULL,
    sage NUMBER(2),
    sdept VARCHAR2(20)
);

CREATE Table course (
    cno NUMBER(1) PRIMARY KEY,
    cname VARCHAR2(30) NOT NULL,
    cpno NUMBER(1),
    credit NUMBER(1) NOT NULL
);

CREATE Table sc (
    sno NUMBER(5),
    cno NUMBER(1),
    grade INT,
    CONSTRAINT foreign1 FOREIGN KEY (sno) REFERENCES student(sno),
    CONSTRAINT foreign2 FOREIGN KEY (cno) REFERENCES course(cno),
    PRIMARY KEY(sno, cno)
);

-- 2. 插入数据
-- INSERT INTO
--     student(sno, sname, sex, sage, sdept)
-- VALUES
--     (95001, '李勇', '男', 19, 'SC');
-- INSERT INTO
--     student(sno, sname, sex, sage, sdept)
-- VALUES(95002, '刘晨', '女', 19, 'IS');
-- INSERT INTO
--     student(sno, sname, sex, sage, sdept)
-- VALUES
--     (95003, '张名', '女', 18, 'MA');
-- INSERT INTO
--     student(sno, sname, sex, sage, sdept)
-- VALUES
--     (95004, '张立', '男', 20, 'IS');
INSERT
    ALL INTO student(sno, sname, sex, sage, sdept)
VALUES
    (95001, '李勇', '男', 19, 'SC') INTO student(sno, sname, sex, sage, sdept)
VALUES
    (95002, '刘晨', '女', 19, 'IS') INTO student(sno, sname, sex, sage, sdept)
VALUES
    (95003, '张名', '女', 18, 'MA') INTO student(sno, sname, sex, sage, sdept)
VALUES
    (95004, '张立', '男', 20, 'IS')
SELECT
    1
FROM
    dual;

-- INSERT INTO course(cno, cname, cpnp, credt) VALUES (1, '数据库', 5, 4);
-- INSERT INTO
--     course(cno, cname, cpnp, credt)
-- VALUES
--     (2, '数学', NULL, 2);
-- INSERT INTO course(cno, cname, cpnp, credt) VALUES (3, '信息系统', 1, 4);
-- INSERT INTO course(cno, cname, cpnp, credt) VALUES (4, '操作系统', 6, 3);
-- INSERT INTO course(cno, cname, cpnp, credt) VALUES (5, '数据结构', 7, 4);
-- INSERT INTO
--     course(cno, cname, cpnp, credt)
-- VALUES
--     (6, '数据处理', NULL, 2);
-- INSERT INTO
--     course(cno, cname, cpnp, credt)
-- VALUES
--     (7, 'JAVA语言', 6, 4);
INSERT
    ALL INTO course(cno, cname, cpno, credit)
VALUES
    (1, '数据库', 5, 4) INTO course(cno, cname, cpno, credit)
VALUES
    (2, '数学', NULL, 2) INTO course(cno, cname, cpno, credit)
VALUES
    (3, '信息系统', 1, 4) INTO course(cno, cname, cpno, credit)
VALUES
    (4, '操作系统', 6, 3) INTO course(cno, cname, cpno, credit)
VALUES
    (5, '数据结构', 7, 4) INTO course(cno, cname, cpno, credit)
VALUES
    (6, '数据处理', NULL, 2) INTO course(cno, cname, cpno, credit)
VALUES
    (7, 'JAVA语言', 6, 4)
SELECT
    1
FROM
    dual;

INSERT
    ALL INTO sc(sno, cno, grade)
VALUES
    (95001, 1, 92) INTO sc(sno, cno, grade)
VALUES
    (95001, 2, 85) INTO sc(sno, cno, grade)
VALUES
    (95001, 5, 88) INTO sc(sno, cno, grade)
VALUES
    (95002, 2, 90) INTO sc(sno, cno, grade)
VALUES
    (95002, 3, 95) INTO sc(sno, cno, grade)
VALUES
    (95003, 2, 80) INTO sc(sno, cno, grade)
VALUES
    (95004, 4, 90)
SELECT
    1
FROM
    dual;

-- （1） 查看表 sc 的结构定义
DESC sc;

-- （2） 修改 student 表结构将 dept 定为 VARCHAR（20）
ALTER TABLE
    student
MODIFY
    sdept VARCHAR(20);

-- （3） 为 student 增加一列 TEL
ALTER TABLE
    student
ADD
    telephone NUMBER(11);

-- （4） 为 student 表的 sex 列添加 check 约束，要求只能传递“男”或“女”
ALTER TABLE
    student
ADD
    CHECK(sex in ('男', '女'));

INSERT INTO
    student(sno, sex)
VALUES
    (12, '其');

-- （5） 为 course 表的 CNAME 列添加 unique 约束
ALTER TABLE
    course
ADD
    UNIQUE(cname);

-- （6） 在 student 表的 SNAME 列上建立索引，索引名为 index_name
CREATE INDEX index_name ON student(sname);

-- （7） 建立信息系（IS）男生的视图 view_cs,包括学号（sno），姓名（sname），年龄（sage）
DROP VIEW view_cs;

CREATE VIEW view_cs AS
SELECT
    sno,
    sname,
    sage
FROM
    student s
WHERE
    s.sdept = 'IS'
    AND s.sex = '男';

SELECT
    *
FROM
    view_cs;

-- （8） 建立学生选课情况视图 SC_view,包括学号，姓名，课程名（cname），学分（credit）
-- 成绩（grade）
CREATE VIEW SC_view AS
SELECT
    s.sno,
    s.sname,
    c.cname,
    c.credit,
    sc.grade
FROM
    student s,
    course c,
    sc sc
WHERE
    s.sno = sc.sno
    AND sc.cno = c.cno;

SELECT
    *
FROM
    SC_view;

-- （9） 创建 course 表的共有同义词 c
CREATE synonym c for course;

-- （10） 查询学生总人数
SELECT
    count(*)
FROM
    student;

-- （11） 查询 2 号课程平均成绩及总成绩
SELECT
    SUM(sc.grade) sum,
    AVG(sc.grade) avg
FROM
    sc
WHERE
    sc.cno = 2;

-- （12） 查询各个课程号与相应的选课人数
SELECT
    sc.cno,
    count(*)
FROM
    sc
GROUP BY
    sc.cno;

-- （13） 查询信息系（IS）学生和性别为女性的学生学号及姓名
SELECT
    sno,
    sname
FROM
    student
WHERE
    sdept = 'IS'
    AND sex = '女';

-- （14） 查询信息系（IS）平均成绩大于 90 分的学生学号
SELECT
    sno
FROM
    sc
GROUP BY
    sc.sno
HAVING
    AVG(sc.grade) > 90;

-- （15） 查询与“刘晨”在同一个系学习的学生
SELECT
    *
FROM
    student s
where
    s.sdept = (
        SELECT
            sdept
        FROM
            student s
        WHERE
            s.sname = '刘晨'
    );

-- （16） 查询选修了课程名为’数学’的学生学号和姓名
SELECT
    s.sno,
    s.sname
FROM
    student s,
    (
        SELECT
            *
        FROM
            sc
        WHERE
            sc.cno = (
                SELECT
                    cno
                FROM
                    course c
                WHERE
                    c.cname = '数学'
            )
    ) sc
WHERE
    s.sno = sc.sno;

-- （17） 找出选修课程中成绩最高的同学的姓名及成绩
SELECT
    s.sno,
    sc.grade
FROM
    student s,
    (
        SELECT
            sno,
            grade
        FROM
            (
                SELECT
                    *
                FROM
                    sc
                ORDER BY
                    grade
            )
        WHERE
            rownum <= 1
    ) sc
WHERE
    s.sno = sc.sno;

-- （18） 查询信息系（IS）选修了 2 门及以上（>=2）课程的学生的学号
SELECT
    s.sno
FROM
    student s,
    (
        SELECT
            *
        FROM
            (
                SELECT
                    sno,
                    COUNT(*) cnt
                FROM
                    sc
                GROUP BY
                    sc.sno
            ) sc
        WHERE
            sc.cnt >= 2
    ) sc
WHERE
    s.sno = sc.sno
    AND s.sdept = 'IS';

-- （19） 查询所有选修了 3 号课程的学生姓名
SELECT
    sname
FROM
    student s,
(
        SELECT
            sno
        FROM
            sc
        WHERE
            sc.cno = 3
    ) sc
WHERE
    s.sno = sc.sno;

-- （20） 查询所有课程的选修情况，结果显示课程号，课程名，学号及成绩（结果中包含未
-- 被选修的课程）
SELECT
    c.cno,
    c.cname,
    s.sno,
    sc.grade
FROM
    student s,
    course c,
    sc
WHERE
    s.sno = sc.sno
    AND sc.cno = c.cno;