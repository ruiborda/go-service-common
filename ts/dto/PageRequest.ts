import { PageRequestOrder } from './PageRequestOrder';
import { PageRequestFilter } from './PageRequestFilter';

export interface PageRequest {
  pageNumber?: number;
  pageSize?: number;
  search?: string;
  sort?: PageRequestOrder[];
  filters?: PageRequestFilter[];
}
