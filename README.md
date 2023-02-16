# adminx

基于golang &amp; Ant Design的企业级权限管理系统

## Ent

```bash
woco entimport --dsn root:@tcp(localhost:3306)/portal --tables User
```

### 迁移

安装atlas
```bash
curl -sSf https://atlasgo.sh | sh
```

### 制作迁移文件

```bash
make mgr DSN=mysql://root@localhost:3306/portal NAME=init
```
> 由于实际采用的参数 autoreplay 需要空数据库,因此dsn可以指向其他空库,不一定需要portal库

### 验证迁移文件
```bash
atlas migrate lint \
--dev-url="mysql://root@localhost:3306/portal" \
--dir="file://ent/migrate/migrations" \
--latest=1
```

### 执行迁移

```bash
atlas migrate apply \
  --dir "file://ent/migrate/migrations" \
  --url mysql://root@localhost:3306/portal
```

cashbin需要手动创建

```shell
atlas migrate new cashbin \
  --dir "file://ent/migrate/migrations"
```
写入以下内容
```
-- create "casbin_rules" table
CREATE TABLE `casbin_rules` (`id` bigint NOT NULL AUTO_INCREMENT, `ptype` varchar(255) NOT NULL DEFAULT '', `v0` varchar(255) NOT NULL DEFAULT '', `v1` varchar(255) NOT NULL DEFAULT '', `v2` varchar(255) NOT NULL DEFAULT '', `v3` varchar(255) NOT NULL DEFAULT '', `v4` varchar(255) NOT NULL DEFAULT '', `v5` varchar(255) NOT NULL DEFAULT '', PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
```

## 数字字典定义

### 常规约定

status: 常规状态,如果涉及工作流状态,请根据业务定义

- 80A : 启用
- 8FF : 禁用

### DEMO

用于测试的JWT("admin secret"):
```
eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiI2N2E4NzQ4MmU5MWY0ZjJlOTIyMGY1MTM3NjE4NWI3ZSIsInN1YiI6IjEiLCJuYW1lIjoiYWRtaW4iLCJpYXQiOjE1MTYyMzkwMjJ9.MzfOsaGAtHj0IIoVDgpfUD1GMtgLTNbY7D_CkEoR9CQ
```