import path from "path";
import type { CodegenConfig } from "@graphql-codegen/cli";

const config: CodegenConfig = {
  overwrite: true,
  schema: [
    path.resolve(__dirname, "../graph/schema/schema.graphql"),
    path.resolve(__dirname, "../graph/schema/types/*.graphql"),
  ],
  // documents: "src/**/*.vue",
  generates: {
    "src/gql/": {
      preset: "client",
      plugins: [],
    },
    "./graphql.schema.json": {
      plugins: ["introspection"],
    },
  },
};

export default config;
