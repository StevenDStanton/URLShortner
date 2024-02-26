import { incrementBase68String } from './base68';
describe('incrementBase68String', () => {
  it('increments strings correctly', () => {
    expect(incrementBase68String('A')).toEqual('B');
    expect(incrementBase68String('Z')).toEqual('-');
    expect(incrementBase68String('9')).toEqual('a');
    expect(incrementBase68String('123ASb!')).toEqual('123ASc0');
  });
});
