import { FilterOperator } from './FilterOperator';

export type FilterValuePrimitive = string | number | boolean | null;
export type FilterValue = FilterValuePrimitive | FilterValuePrimitive[];

export interface PageRequestFilter {
  field: string;
  operator: FilterOperator;
  value: FilterValue;
}
