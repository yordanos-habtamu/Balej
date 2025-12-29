import { createClient, fetchExchange, cacheExchange } from 'urql';
import { useCookie, useNuxtApp } from 'nuxt/app';

export const useUrqlClient = () => {
  const { ssrContext } = useNuxtApp();
  const isServer = !!ssrContext;

  let token = null;
  if (isServer) {
    token = useCookie('token'); // Only call useCookie on server
  } else if (process.client) {
    // Client-side: Access cookies via document.cookie or other methods
    token = document.cookie
      .split('; ')
      .find(row => row.startsWith('token='))
      ?.split('=')[1] || null;
  }

  const client = createClient({
    url: 'http://localhost:7000/graphql',
    fetchOptions: {
      credentials: 'include',
      headers: token ? { Authorization: `Bearer ${token}` } : {},
    },
    exchanges: [cacheExchange, fetchExchange],
  });

  return client;
};