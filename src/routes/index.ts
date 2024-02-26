import { Router, Request, Response } from 'express';

const router: Router = Router();

router.get('/', (req: Request, res: Response) => {
  res.send({ status: 'I Am Alive!' });
});

export { router as indexRouter };
