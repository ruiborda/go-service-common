import { PageRequest } from '../PageRequest';

export interface GetOrdersPaginatedRequest extends PageRequest {
  startDate?: string;
  endDate?: string;
}
