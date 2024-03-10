import { createClient } from '@libsql/client';

const client = createClient({
  url: 'libsql://wxs-us-stevendstanton.turso.io',
  authToken: '389f22c4-eded-4354-9bca-268a58eb6ade',
});

// await client.execute({
//   sql: "INSERT INTO users VALUES (:name)",
//   args: { name: "Iku" },
// });

export async function getURL(indexKey: string): Promise<string> {
  const results = await client.execute({
    sql: 'SELECT url from url_map WHERE index_key = (:indexKey)',
    args: { indexKey: indexKey },
  });
  if (
    results &&
    results.rows.length > 0 &&
    typeof results.rows[0].url === 'string'
  ) {
    return results.rows[0].url;
  }
}

export async function putURL(indexKey: string, url: string): Promise<boolean> {
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
  const row = await db.get('SELECT value FROM latest_record WHERE id = 0');
  await db.close();
  return row ? row.value : null;
}

export async function setLatestIndex(value: string): Promise<boolean> {
  try {
    await db.run('UPDATE latest_record SET value = ? WHERE id = 0', value);
    await db.close();
    return true;
  } catch (error) {
    console.error('Error in setLatestIndex:', error);
    await db.close();
    return false;
  }
}
