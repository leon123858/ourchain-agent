# API 文件

詳情請參考以下兩個網址
- [javascript](https://github.com/leon123858/ourChain-frontend/tree/main/ourchain-web-cli)
- [flutter](https://github.com/leon123858/ourChain-frontend/tree/main/our-wallet-app)

### 基本資訊

**基礎 URL:** 依您的服務器配置而定

**共通響應結構:**

成功時返回:
```json
{
  "result": "success",
  "data": {}
}
```

失敗時返回:
```json
{
  "result": "fail",
  "error": "錯誤訊息"
}
```

---

### GET 請求

#### 獲取餘額
- **URL:** `/get/balance`
- **方法:** `GET`
- **參數:**
    - `address`: string (必須)
- **響應:** 餘額資訊

#### 獲取私鑰
- **URL:** `/get/privatekey`
- **方法:** `GET`
- **參數:**
    - `address`: string (必須)
- **響應:** 私鑰資訊

#### 獲取交易資訊
- **URL:** `/get/transaction`
- **方法:** `GET`
- **參數:**
    - `txid`: string (必須)
- **響應:** 交易資訊

#### 獲取未花費輸出 (UTXO)
- **URL:** `/get/utxo`
- **方法:** `GET`
- **參數:**
    - `address`: string (非必須)
- **響應:** UTXO 列表

#### 獲取合約資訊
- **URL:** `/get/contract`
- **方法:** `GET`
- **參數:**
    - `protocol`: string (非必須)
- **響應:** 合約列表

#### 生成新地址
- **URL:** `/get/newaddress`
- **方法:** `GET`
- **參數:** 無
- **響應:** 新地址

---

### POST 請求

note: 以下參數小寫開頭為 query string 參數，大寫開頭為 body 參數，實際傳輸依舊用小寫開頭

#### 生成區塊
- **URL:** `/block/generate`
- **方法:** `POST`
- **參數:**
    - `address`: string (非必須)
- **響應:** 生成的區塊 ID 列表

#### 傾印合約訊息
- **URL:** `/get/contractmessage`
- **方法:** `POST`
- **參數:**
    - `Address`: string (必須)
    - `Arguments`: string (必須)
- **響應:** 合約訊息

```javascript
const result = await fetch(`${BASE_URL}get/contractmessage`, {
    method: 'POST',
    headers: {
        'Content-Type': 'application/json',
    },
    body: JSON.stringify({
        address: targetAddress,
        arguments: args,
    }),
});
```

#### 創建原始交易
- **URL:** `/rawtransaction/create`
- **方法:** `POST`
- **參數:**
    - `Inputs`: array (必須)
    - `Outputs`: object (必須)
    - `Contract`: string (非必須)
- **響應:** 原始交易資訊

```javascript
const result = await fetch(`${BASE_URL}rawtransaction/create`, {
    method: 'POST',
    headers: {
        'Content-Type': 'application/json',
    },
    body: JSON.stringify({
        inputs: [utxo.input],
        outputs: [utxo.output],
        contract: {
            action: contract.action,
            code: contract.code,
            address: contract.address,
            args: contract.args,
        },
    }),
});
```

#### 發送原始交易
- **URL:** `/rawtransaction/send`
- **方法:** `POST`
- **參數:**
    - `RawTransaction`: string (必須)
- **響應:** 發送結果

```javascript
const result = await fetch(`${BASE_URL}rawtransaction/send`, {
    method: 'POST',
    headers: {
        'Content-Type': 'application/json',
    },
    body: JSON.stringify({
        rawTransaction: signedTx,
    }),
});
```
#### 簽名原始交易
- **URL:** `/rawtransaction/sign`
- **方法:** `POST`
- **參數:**
    - `RawTransaction`: string (必須)
    - `PrivateKey`: string (必須)
- **響應:** 簽名結果

```javascript
const result = await fetch(`${BASE_URL}rawtransaction/sign`, {
    method: 'POST',
    headers: {
        'Content-Type': 'application/json',
    },
    body: JSON.stringify({
        rawTransaction: rawTx,
        privateKey: privateKey,
    }),
});
```

---



### 封裝範例

[以 javascript 為例](https://github.com/leon123858/ourChain-frontend/blob/main/ourchain-web-cli/src/utils/txApiWrapper.ts)

[以 flutter 為例](https://github.com/leon123858/ourChain-frontend/tree/main/our-wallet-app/lib/services/chain)

概念解釋:

所有交易的創建都有三個流程:
1. 創建原始交易
2. 簽名原始交易
3. 發送原始交易

所以需要按照順序調用這三個 post 方法, 並且將上一個方法的結果作為下一個方法的參數, 另外為了創建交易所需要的資金(utxo)需要先獲取 utxo 作為創建交易的 input

以上流程可以創建一個交易, 基於交易有以下兩種合約操作
- 部署合約: 在區塊鏈上部署智能合約, 並且返回合約地址和交易編號
- 執行合約: 在區塊鏈上執行智能合約, 返回交易編號

最後, 若要讀取數據可以直接調用以下方法, 此方法無需基於交易
- 傾印合約訊息: 獲取合約執行結果

結論, 智能合約的使用可以被視為是[非同步後端調用](https://ithelp.ithome.com.tw/articles/10338803), 概念上就是利用 deployContract 非同步定義要做的行為, 利用 callContract 非同步執行合約。
最後用 getContractMessage 同步獲取合約執行結果。

### scanner API 說明

以下 API 皆為 GET 請求, 作為 scanner 被呼叫
- `/get/utxo` 獲取未花費輸出 (UTXO)
- `/get/contract` 獲取合約資訊

基於 bitcoin 的 utxo 集沒有使用 address 作為 index, 所以需要掃描所有區塊, 並且過濾出符合條件的 utxo
, 這個過程是非常耗時的, 所以需要一個 scanner 來完成這個任務, 並且將結果存入資料庫, 以供後續使用。常見的使用情境是創建交易時需要解鎖 utxo。

合約資訊的獲取是基於區塊鏈上的合約"一般化介面", 這個資訊是不會變動的, 所以只需要掃描一次即可, 並且將結果存入資料庫, 以供後續使用。常見的使用情境是針對特定協議的合約產生列表。