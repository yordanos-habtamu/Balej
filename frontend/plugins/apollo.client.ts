import { defineNuxtPlugin } from '#app'
import { ApolloClient, InMemoryCache, createHttpLink } from '@apollo/client/core'
import { DefaultApolloClient } from '@vue/apollo-composable'

export default defineNuxtPlugin((nuxtApp) => {
  const httpLink = createHttpLink({
    uri: 'http://localhost:8080/graphql', // your Go backend GraphQL endpoint
  })

  const apolloClient = new ApolloClient({
    link: httpLink,
    cache: new InMemoryCache(),
  })

  nuxtApp.vueApp.provide(DefaultApolloClient, apolloClient)
})
