import {
  getURL,
  putURL,
  getLatestIndex,
  setLatestIndex,
  initializeDb,
} from './sqliteDb';

beforeEach(async () => {
  process.env.NODE_ENV = 'test';
  await initializeDb();
});

afterEach(async () => {});

describe('Database functions', () => {
  test('putURL inserts a new URL correctly', async () => {
    const result = await putURL('testKey', 'http://example.com');
    expect(result).toBe(true);

    const url = await getURL('testKey');
    expect(url).toBe('http://example.com');
  });

  test('getLatestIndex returns the correct latest index', async () => {
    await setLatestIndex('testIndex');
    const index = await getLatestIndex();
    expect(index).toBe('testIndex');
  });
});
