import sqlite3 from 'sqlite3';
import { Database, open } from 'sqlite';

export async function openDb() {
  return open({
    filename: './db.sqlite',
    driver: sqlite3.Database,
  });
}

export async function getURL(indexKey: string): Promise<string> {
  const db = await openDb();
  const row = await db.get('SELECT url FROM url_map WHERE index_key = ?', [
    indexKey,
  ]);
  await db.close();
  return row ? row.url : null;
}

export async function putURL(indexKey: string, url: string): Promise<boolean> {
  const db = await openDb();
  try {
    await db.run(
      `INSERT INTO url_map (index_key, url) VALUES (?, ?)
            ON CONFLICT(index_key) DO UPDATE SET url = excluded.url`,
      [indexKey, url],
    );
    await db.close();
    return true;
  } catch (error) {
    console.error('Error in putURL:', error);
    await db.close();
    return false;
  }
}

export async function getLatestIndex(): Promise<string> {
  const db = await openDb();
  const row = await db.get('SELECT value FROM latest_record WHERE id = 0');
  await db.close();
  return row
    ? Buffer.from(row.value, 'base64').toString()
    : Buffer.from('0').toString('base64');
}

export async function setLatestIndex(value: string): Promise<boolean> {
  const db = await openDb();
  try {
    const base64Value = Buffer.from(value).toString('base64');
    await db.run(
      'UPDATE latest_record SET value = ? WHERE id = 0',
      base64Value,
    );
    await db.close();
    return true;
  } catch (error) {
    console.error('Error in setLatestIndex:', error);
    await db.close();
    return false;
  }
}

export async function initializeDb() {
  const db = await openDb();
  await db.exec(`
      CREATE TABLE IF NOT EXISTS latest_record (
          id INTEGER PRIMARY KEY,
          value TEXT
      )
    `);
  await db.run(`
      INSERT INTO latest_record(id, value)
      VALUES(0, '0')
      ON CONFLICT(id) DO NOTHING
    `);
  await db.exec(`
      CREATE TABLE IF NOT EXISTS url_map (
          index_key TEXT PRIMARY KEY,
          url TEXT
      )
    `);

  await db.close();
}
