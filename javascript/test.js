const http = require('./http');
const creator = require('./random');
const log = require('tracer').colorConsole({
  format: '{{message}}',
});

// eth_getBalance
// eth_getTransactionCount
// Random Test
// 1. Data
// 2. Quantity|Tag
exports.dataQuantityTagRandom = (
  method, // string
  dataType, // hex, string
  dataByte, // number
  quantityTagType, // quantity, tag
  maxQuantity, // number
  testCount // number
) => {
  for (let i = 0; i < testCount; i++) {
    let data, quantityOrTag, params;
    switch (dataType) {
      case 'hex':
        data = creator.randomHex(2 * dataByte);
        break;
      case 'string':
        data = creator.randomString(dataByte);
        break;
      default:
        log.error('파라미터 입력값을 확인해주세요.');
    }
    switch (quantityTagType) {
      case 'quantity':
        quantityOrTag = creator.randomQuantity(maxQuantity);
        break;
      case 'tag':
        quantityOrTag = creator.randomTag();
        break;
      default:
        log.error('파라미터 입력값을 확인해주세요.');
    }
    params = [data, quantityOrTag];
    http.post(method, params);
  }
};

// eth_getBalance
// eth_getTransactionCount
// Test
// 1. Data
// 2. Quantity|Tag
exports.dataQuantityTag = (
  method, // string
  data,
  quantityTag
) => {
  params = [data, quantityTag];
  http.post(method, params);
};
