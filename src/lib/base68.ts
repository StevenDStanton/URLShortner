const BASE68_CHARS =
  '0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_*+!';
const BASE_LENGTH = BASE68_CHARS.length;

export function incrementBase68String(str: String) {
  if (str === '') return BASE68_CHARS[0];

  let indices = str.split('').map((char) => BASE68_CHARS.indexOf(char));
  let carry = 1;

  for (let i = indices.length - 1; i >= 0; i--) {
    if (carry === 0) break;

    let newVal = indices[i] + carry;
    carry = newVal >= BASE_LENGTH ? 1 : 0;
    indices[i] = newVal % BASE_LENGTH;
  }

  if (carry > 0) {
    indices.unshift(0);
  }

  return indices.map((i) => BASE68_CHARS[i]).join('');
}
