const characters =
  'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
const hexCharacters = 'abcdef0123456789';
const tags = ['latest', 'earliest', 'pending'];

exports.randomString = (size) => {
  let result = '';
  for (let i = 0; i < size; i++) {
    result += characters.charAt(Math.floor(Math.random() * characters.length));
  }
  return result;
};

exports.randomHex = (size) => {
  let result = '0x';
  for (let i = 0; i < size; i++) {
    result += hexCharacters.charAt(
      Math.floor(Math.random() * hexCharacters.length)
    );
  }
  return result;
};

exports.randomQuantity = (max) => {
  let random = Math.floor(Math.random() * (max + 1));
  let hex = random.toString(16);
  let result = '0x' + hex;
  return result;
};

exports.randomBoolean = () => {
  if (Math.floor(Math.random() * 2)) return true;
  return false;
};

exports.randomTag = () => {
  return tags[Math.floor(Math.random() * tags.length)];
};
