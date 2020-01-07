# 前后端交互 API

#### /api/metadata

后端元信息

```json
{
	"static": false
}
```

#### /api/problem-set-list

返回题库列表

```json
[
	{"name": "LOJ","id": "loj"},
	...
	{"name":"UOJ", "id": "uoj"}
]
```

#### /api/problem-set/\<problemset\>/metadata

 返回指定题库的信息

* problemset: 题库id



 ```json
 {
	"name": "LOJ",
	"problem": 1000,
	"page": "50"
 }
 ```

#### /api/problem-set/\<problemset\>/css

返回指定题库的css数据

* problemset: 题库id

 #### /api/problem-list/\<problemset\>/\<page\>

 返回题库 \<problemset\> 第 \<page\> 页的题目

 * problemset: 题库id
 * page: 页码



```json
[
	{"pid": "1", "title": "A + B Problem"},
	...
	{"pid": "101", "title": "最大流"}
]
```

#### /api/problem/\<problemset\>/\<problem\>

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
    "description": "<div class=\"oiarchive-block\"></div>\n<div class=\"oiarchive-block\">\n<h4 class=\"oiarchive-block-title\">题目描述</h4>\n\n<p>输入 $ a $ 和 $ b $，输出 $ a + b $ 的结果。</p>\n\n</div>\n<div class=\"oiarchive-block\">\n<h4 class=\"oiarchive-block-title\">输入格式</h4>\n\n<p>一行两个正整数 $ a $ 和 $ b $。</p>\n\n</div>\n<div class=\"oiarchive-block\">\n<h4 class=\"oiarchive-block-title\">输出格式</h4>\n\n<p>一行一个正整数 $ a + b $。</p>\n\n</div>\n<div class=\"oiarchive-block\">\n<h4 class=\"oiarchive-block-title\">样例</h4>\n\n<h4>样例输入</h4>\n\n<pre><code class=\"language-plain\">1 2\n</code></pre>\n\n<h4>样例输出</h4>\n\n<pre><code class=\"language-plain\">3\n</code></pre>\n\n</div>\n<div class=\"oiarchive-block\">\n<h4 class=\"oiarchive-block-title\">数据范围与提示</h4>\n\n<p>对于 $ 100\\% $ 的数据，$ 1 \\leq a, b \\leq 10 ^ 6 $。</p>\n\n</div>\n"
}
```

#### /api/history-list/\<problemset\>/\<problem\>/\<page\>

返回历史版本列表

* page : 页码数

#### /api/history/\<problemset\>/\<problem\>/\<commit-id\>/

返回 \<commit-id\> 版本时的题面数据


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