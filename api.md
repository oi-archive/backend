# 前后端交互 API



#### /api/problem-set-list

返回题库列表

```json
[
	{"name": "LOJ","id": "loj"},
	...
	{"name":"UOJ", "id": "uoj"}
]
```

#### /api/problem-set/<problemset>/metadata

 返回指定题库的信息

* problemset: 题库id



 ```json
 {
	"name": "LOJ",
	"problem": 1000,
	"page": "50"
 }
 ```

#### /api/problem-set/<problemset>/css

返回指定题库的css数据

* problemset: 题库id

 #### /api/problem-set/<problemset>/page/<page>

 返回题库 <problemset> 第 <page> 页的题目

 * problemset: 题库id
 * page: 页码



```json
[
	{"pid": "1", "name": "A + B Problem"},
	...
	{"pid": "101", "name": "最大流"}
]
```

#### /api/problem/<problemset>/<problem>

返回指定题目信息

* problemset: 题库id
* problem: 题目id

```json
{
	"title": "A + B Problem",
	"time": 2000,
	"memory": 512,
	"judge": "传统",
	"url":"https://loj.ac/problem",
	"description_type": "markdown",
	"description": "# 题目描述\n\n输入 $ a $ 和 $ b $，输出 $ a + b $ 的结果。\n\n# 输入格式\n\n一行两个正整数 $ a $ 和 $ b $。"
}
```



#### /api/search

**POST请求**

Post form:
```json
{
	"problemset": "loj",
	"data": "A + B"
}
```
注： problemset 项若为空则表示在所有题库进行搜索

Response:
```json
[
	{"problemset": "LOJ", "pid": "1", "name": "A + B Problem"},
	...
	{"problemset": "UOJ", "pid": "1", "name": "A + B Problem"}
]
```