const test = require('./test');

test.dataQuantityTagRandom('eth_getBalance', 'hex', 20, 'tag', 1, 5);
test.dataQuantityTag(
  'eth_getBalance',
  '0xabcf409adf429c4d925a0300ce9de34a870afff0',
  'latest'
);

test.dataQuantityTagRandom('eth_getTransactionCount', 'hex', 20, 'tag', 1, 5);
test.dataQuantityTag(
  'eth_getTransactionCount',
  '0xabcf409adf429c4d925a0300ce9de34a870afff0',
  'latest'
);
