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

// eth_call
// eth_estimateGas
// Random Test
// 1. Object
// 2. Quantity|Tag
exports.objectQuantityTagRandom = (
  method, // string
  fromType, // hex, string, null
  fromByte, // number
  toType, // hex, string, null
  toByte, // number,
  maxGas, // number, null
  maxGasPrice, // number, null
  maxValue, // number, null
  dataType, // hex, string, null
  dataByte, // number
  quantityTagType, // quantity, tag, null
  maxQuantity, // number
  testCount // number
) => {
  for (let i = 0; i < testCount; i++) {
    let from, to, gas, gasPrice, value, data, quantityOrTag, params;
    switch (fromType) {
      case 'hex':
        from = creator.randomHex(2 * fromByte);
        break;
      case 'string':
        from = creator.randomString(fromByte);
        break;
      case null:
        from = null;
        break;
      default:
        log.error('파라미터 입력값을 확인해주세요.');
    }
    switch (toType) {
      case 'hex':
        to = creator.randomHex(2 * toByte);
        break;
      case 'string':
        to = creator.randomString(toByte);
        break;
      case null:
        to = null;
        break;
      default:
        log.error('파라미터 입력값을 확인해주세요.');
    }
    if (maxGas == null) {
      gas = null;
    } else {
      gas = creator.randomQuantity(maxGas);
    }
    if (maxGasPrice == null) {
      gasPrice = null;
    } else {
      gasPrice = creator.randomQuantity(maxGasPrice);
    }
    if (maxValue == null) {
      value = null;
    } else {
      value = creator.randomQuantity(maxValue);
    }
    switch (dataType) {
      case 'hex':
        data = creator.randomHex(2 * dataByte);
        break;
      case 'string':
        data = creator.randomString(dataByte);
        break;
      case null:
        data = null;
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
      case null:
        quantityOrTag = null;
        break;
      default:
        log.error('파라미터 입력값을 확인해주세요.');
    }
    params = [
      {
        from: from,
        to: to,
        gas: gas,
        gasPrice: gasPrice,
        value: value,
        data: data,
      },
      quantityOrTag,
    ];
    http.post(method, params);
  }
};

// eth_call
// eth_estimateGas
// Test
// 1. Object
// 2. Quantity|Tag
exports.objectQuantityTag = (
  method, // string
  from,
  to,
  gas,
  gasPrice,
  value,
  data,
  quantityTag
) => {
  params = [
    {
      from: from,
      to: to,
      gas: gas,
      gasPrice: gasPrice,
      value: value,
      data: data,
    },
    quantityTag,
  ];
  http.post(method, params);
};

// eth_getBlockByHash
// Random Test
// 1. Data
// 2. Boolean
exports.dataBooleanRandom = (
  method, // string
  dataType, // hex, string
  dataByte, // number
  testCount // number
) => {
  for (let i = 0; i < testCount; i++) {
    let data, boolean, params;
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
    boolean = creator.randomBoolean();
    params = [data, boolean];
    http.post(method, params);
  }
};

// eth_getBlockByHash
// Test
// 1. Data
// 2. Boolean
exports.dataBoolean = (
  method, // string
  data,
  boolean
) => {
  params = [data, boolean];
  http.post(method, params);
};

// eth_getBlockByNumber
// Random Test
// 1. Quantity|Tag
// 2. Boolean
exports.quantityTagBooleanRandom = (
  method, // string
  quantityTagType, // quantity, tag
  maxQuantity, // number
  testCount // number
) => {
  for (let i = 0; i < testCount; i++) {
    let quantityOrTag, boolean, params;
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
    boolean = creator.randomBoolean();
    params = [quantityOrTag, boolean];
    http.post(method, params);
  }
};

// eth_getBlockByNumber
// Test
// 1. Quantity|Tag
// 2. Boolean
exports.quantityTagBoolean = (
  method, // string
  quantityTag,
  boolean
) => {
  params = [quantityTag, boolean];
  http.post(method, params);
};

// eth_getTransactionByHash
// eth_getTransactionReceipt
// Random Test
// 1. Data
exports.dataRandom = (
  method, // string
  dataType, // hex, string
  dataByte, // number
  testCount // number
) => {
  for (let i = 0; i < testCount; i++) {
    let data, params;
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
    params = [data];
    http.post(method, params);
  }
};

// eth_getTransactionByHash
// eth_getTransactionReceipt
// Test
// 1. Data
exports.data = (
  method, // string
  data
) => {
  params = [data];
  http.post(method, params);
};
