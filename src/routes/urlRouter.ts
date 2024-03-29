import { Router, Request, Response } from 'express';
import {
  getURL,
  putURL,
  getLatestIndex,
  setLatestIndex,
} from '../lib/sqliteDb';
import { incrementBase68String } from '../lib/base68';

const router: Router = Router();

router.put('/:index?', async (req: Request, res: Response) => {
  try {
    const { url } = req.body;
    if (!url) {
      return res.status(400).send('URL is required');
    }
    let genKey = false;
    let indexKey = req.params.index;
    if (!indexKey) {
      genKey = true;
      indexKey = await getLatestIndex();
    }
    const success = await putURL(indexKey, url);

    if (success) {
      if (genKey) {
        const newIndexKey = incrementBase68String(indexKey);
        await setLatestIndex(newIndexKey);
      }
      const message = `URL associated with index: ${indexKey} has been updated.`;
      res.json({ index: indexKey, url });
    } else {
      res.status(500).send('Failed to update the database');
    }
  } catch (error) {
    res.status(500).send('Internal Server Error');
  }
});

router.get('/:index', async (req: Request, res: Response) => {
  try {
    const indexKey = req.params.index;
    const url = await getURL(indexKey);
    if (url) {
      res.redirect(url);
    } else {
      res.status(404).send('URL not found');
    }
  } catch (error) {
    console.error('Error fetching URL:', error);
    res.status(500).send('Internal Server Error');
  }
});

export { router as urlRouter };
