// src/shared/api/baseApi.ts
import { createApi, fetchBaseQuery } from '@reduxjs/toolkit/query/react';

export const baseApi = createApi({
  reducerPath: 'api', // ключ в store
  baseQuery: fetchBaseQuery({
    baseUrl: 'https://example.com/api', // замени на свой реальный API
  }),
  endpoints: () => ({}), // сюда будут добавляться конкретные endpoints
});
