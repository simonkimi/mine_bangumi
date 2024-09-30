import { ApolloClient, createHttpLink, InMemoryCache } from "@apollo/client";
import { setContext } from "@apollo/client/link/context";
import { useUserStore } from "@/stores/userStore";

const httpLink = createHttpLink({
  uri: "/api/v1/graph",
});

const authLink = setContext((_, { headers }) => {
  const userStore = useUserStore();
  const token = userStore.apiToken;
  return {
    headers: {
      ...headers,
      authorization: token ? `Token ${token}` : "",
    },
  };
});

const apolloClient = new ApolloClient({
  cache: new InMemoryCache(),
  link: authLink.concat(httpLink),
});

export default apolloClient;
