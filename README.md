# 智能代码分析平台接口设计文档

---

## 用户认证接口

### 1. 注册用户

**HTTP方法**：`POST`

**URI**：`/api/v1/auth/register`

**请求内容类型**：`application/json`

**请求参数**:

| 参数名   | 类型   | 必填 | 描述     |
|----------|--------|------|----------|
| username | string | 是   | 用户名   |
| password | string | 是   | 密码     |
| email    | string | 是   | 邮箱地址 |

**请求示例**:

```json
{
  "username": "newuser",
  "password": "password123",
  "email": "newuser@example.com"
}
```

**成功响应**：`201 Created`

**响应内容类型**：`application/json`

**响应参数**:

| 参数名   | 类型   | 描述               |
|----------|--------|--------------------|
| userId   | string | 新注册用户的唯一ID |
| message  | string | 操作结果信息       |

**响应示例**:

```json
{
  "userId": "1a2b3c4d5e6f7g8h9i",
  "message": "User registered successfully."
}
```

**错误状态码**:

| 状态码 | 描述                     |
|--------|--------------------------|
| 400    | Bad Request - 参数错误   |
| 409    | Conflict - 用户名已存在  |

---

### 2. 用户登录

**HTTP方法**：`POST`

**URI**：`/api/v1/auth/login`

**请求内容类型**：`application/json`

**请求参数**:

| 参数名   | 类型   | 必填 | 描述   |
|----------|--------|------|--------|
| username | string | 是   | 用户名 |
| password | string | 是   | 密码   |

**请求示例**:

```json
{
  "username": "existinguser",
  "password": "password123"
}
```

**成功响应**：`200 OK`

**响应内容类型**：`application/json`

**响应参数**:

| 参数名 | 类型   | 描述              |
|--------|--------|-------------------|
| token  | string | 授权令牌          |

**响应示例**:

```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

**错误状态码**:

| 状态码 | 描述                                |
|--------|-------------------------------------|
| 400    | Bad Request - 参数错误              |
| 401    | Unauthorized - 凭证无效或未提供凭证 |

---

### 3. 用户注销

**HTTP方法**：`POST`

**URI**：`/api/v1/auth/logout`

**请求内容类型**：`application/x-www-form-urlencoded`

**请求参数**:

无需请求体，仅需在HTTP头中提供有效的认证Token。

**成功响应**：`204 No Content`

无响应体。

**错误状态码**:

| 状态码 | 描述                                |
|--------|-------------------------------------|
| 401    | Unauthorized - 凭证无效或未提供凭证 |

---

## 用户代码管理接口

### 4. 保存代码

**HTTP方法**：`POST`

**URI**：`/api/v1/code`

**请求内容类型**：`application/json`

**请求参数**:

| 参数名    | 类型   | 必填 | 描述           |
|-----------|--------|------|----------------|
| code      | string | 是   | 代码内容       |
| language  | string | 是   | 代码语言类型   |
| title     | string | 否   | 代码片段的标题 |

**请求示例**:

```json
{
  "code": "package main\n\nimport \"fmt\"\n\nfunc main() {\n    fmt.Println(\"Hello, World!\")\n}",
  "language": "go",
  "title": "Hello World Example"
}
```

**成功响应**：`201 Created`

**响应内容类型**：`application/json`

**响应参数**:

| 参数名  | 类型   | 描述               |
|---------|--------|--------------------|
| codeId  | string | 创建的代码片段ID   |
| message | string | 操作结果信息       |

**响应示例**:

```json
{
  "codeId": "abcd1234efgh5678",
  "message": "Code saved successfully."
}
```

**错误状态码**:

| 状态码 | 描述                     |
|--------|--------------------------|
| 400    | Bad Request - 参数错误   |
| 401    | Unauthorized - 凭证无效或未提供凭证 |

---

### 5. 更新代码

**HTTP方法**：`PUT`

**URI**：`/api/v1/code/{codeId}`

**请求内容类型**：`application/json`

**请求参数**:

| 参数名    | 类型   | 必填 | 描述             |
|-----------|--------|------|------------------|
| code      | string | 是   | 更新的代码内容   |
| language  | string | 是   | 代码语言类型     |
| title     | string | 否   | 代码片段的标题   |

**请求示例**:

```json
{
  "code": "package main\n\nimport \"fmt\"\n\nfunc main() {\n    fmt.Println(\"Hello, World! Updated\")\n}",
  "language": "go",
  "title": "Updated Hello World Example"
}
```

**成功响应**：`200 OK`

**响应内容类型**：`application/json`

**响应参数**:

| 参数名  | 类型   | 描述            |
|---------|--------|-----------------|
| codeId  | string | 更新的代码片段ID |
| message | string | 操作结果信息    |

**响应示例**:

```json
{
  "codeId": "abcd1234efgh5678",
  "message": "Code updated successfully."
}
```

**错误状态码**:

| 状态码 | 描述                     |
|--------|--------------------------|
| 400    | Bad Request - 参数错误   |
| 401    | Unauthorized - 凭证无效或未提供凭证 |
| 404    | Not Found - 代码片段未找到 |

---

### 6. 删除代码

**HTTP方法**：`DELETE`

**URI**：`/api/v1/code/{codeId}`

**请求内容类型**：无需请求主体。

**成功响应**：`204 No Content`

无响应体。

**错误状态码**:

| 状态码 | 描述                     |
|--------|--------------------------|
| 401    | Unauthorized - 凭证无效或未提供凭证 |
| 404    | Not Found - 代码片段未找到 |

---

### 7. 获取用户代码列表

**HTTP方法**：`GET`

**URI**：`/api/v1/code`

**请求内容类型**：无需请求主体。

**成功响应**：`200 OK`

**响应内容类型**：`application/json`

**响应参数**:

| 参数名  | 类型  | 描述            |
|---------|-------|-----------------|
| codes   | array | 代码片段列表    |
| codeId  | string| 代码片段ID      |
| title   | string| 代码片段标题    |
| language| string| 代码语言类型    |
| code    | string| 代码内容        |
| createdAt| string| 创建时间       |

**响应示例**:

```json
{
  "codes": [
    {
      "codeId": "abcd1234efgh5678",
      "title": "Hello World Example",
      "language": "go",
      "code": "package main\n\nimport \"fmt\"\n\nfunc main() {\n    fmt.Println(\"Hello, World!\")\n}",
      "createdAt": "2023-04-01T12:00:00Z"
    },
    {
      "codeId": "ijkl91011mnop1213",
      "title": "Another Example",
      "language": "go",
      "code": "package main\n\nfunc main() {\n    // another example\n}",
      "createdAt": "2023-04-02T15:30:00Z"
    }
  ]
}
```

**错误状态码**:

| 状态码 | 描述                     |
|--------|--------------------------|
| 401    | Unauthorized - 凭证无效或未提供凭证 |

---

## 代码分析接口

### 8. 提交代码进行分析

**HTTP方法**：`POST`

**URI**：`/api/v1/analysis`

**请求内容类型**：`application/json`

**请求参数**:

| 参数名    | 类型   | 必填 | 描述           |
|-----------|--------|------|----------------|
| code      | string | 是   | 代码内容       |
| language  | string | 是   | 代码语言类型   |

**请求示例**:

```json
{
  "code": "package main\n\nimport \"fmt\"\n\nfunc main() {\n    fmt.Println(\"Hello, World!\")\n}",
  "language": "go"
}
```

**成功响应**：`202 Accepted`

**响应内容类型**：`application/json`

**响应参数**:

| 参数名  | 类型   | 描述               |
|---------|--------|--------------------|
| taskId  | string | 分析任务的唯一ID   |
| message | string | 操作结果信息       |

**响应示例**:

```json
{
  "taskId": "efgh5678ijkl91011",
  "message": "Code analysis started successfully."
}
```

**错误状态码**:

| 状态码 | 描述                     |
|--------|--------------------------|
| 400    | Bad Request - 参数错误   |
| 401    | Unauthorized - 凭证无效或未提供凭证 |

---

### 9. 获取代码分析状态

**HTTP方法**：`GET`

**URI**：`/api/v1/analysis/{taskId}/status`

**请求内容类型**：无需请求主体。

**成功响应**：`200 OK`

**响应内容类型**：`application/json`

**响应参数**:

| 参数名  | 类型   | 描述                 |
|---------|--------|----------------------|
| taskId  | string | 分析任务的唯一ID     |
| status  | string | 分析任务的当前状态   |

**响应示例**:

```json
{
  "taskId": "efgh5678ijkl91011",
  "status": "In Progress"
}
```

**错误状态码**:

| 状态码 | 描述                     |
|--------|--------------------------|
| 400    | Bad Request - 参数错误   |
| 401    | Unauthorized - 凭证无效或未提供凭证 |
| 404    | Not Found - 分析任务未找到 |

---

### 10. 获取代码分析结果

**HTTP方法**：`GET`

**URI**：`/api/v1/analysis/{taskId}`

**请求内容类型**：无需请求主体。

**成功响应**：`200 OK`

**响应内容类型**：`application/json`

**响应参数**:

| 参数名   | 类型   | 描述                 |
|----------|--------|----------------------|
| taskId   | string | 分析任务的唯一ID     |
| results  | object | 分析结果             |
| findings | array  | 具体的分析发现       |
| detail   | string | 分析发现的详细信息   |
| severity | string | 分析发现的严重程度   |

**响应示例**:

```json
{
  "taskId": "efgh5678ijkl91011",
  "results": {
    "findings": [
      {
        "detail": "The 'fmt' package is imported but not used.",
        "severity": "warning"
      },
      {
        "detail": "Function 'main' is not commented.",
        "severity": "info"
      }
    ]
  }
}
```

**错误状态码**:

| 状态码 | 描述                     |
|--------|--------------------------|
| 400    | Bad Request - 参数错误   |
| 401    | Unauthorized - 凭证无效或未提供凭证 |
| 404    | Not Found - 分析任务未找到 |

---
