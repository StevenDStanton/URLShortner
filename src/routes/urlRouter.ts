import { Router, Request, Response } from 'express';

const router: Router = Router();

router.get('/', (req: Request, res: Response) => {});

router.put('/', async (req: Request, res: Response) => {});

export { router as urlRouter };
