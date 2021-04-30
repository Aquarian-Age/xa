## goscp

- 一個簡單的cli

- 基於[sftp](https://github.com/pkg/sftp)

- 現有功能只能傳輸指定目錄下的文件 不能傳輸子目錄

- 編輯`goscp.json`文件執行

**遠程到本地**
```bash
./goscp -p pull
```

**本地到遠程**
```bash
./goscp -p push
```
