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

    let indexKey = req.params.index;
    if (!indexKey) {
      const latestIndex = await getLatestIndex();
      indexKey = incrementBase68String(latestIndex);
      await setLatestIndex(indexKey);
    }
    const success = await putURL(indexKey, url);
    if (success) {
      const message = `URL associated with index: ${indexKey} has been updated.`;
      res.json({ message, index: indexKey, url });
    } else {
      res.status(500).send('Failed to update the database');
    }
  } catch (error) {
    console.error('Error processing URL:', error);
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
