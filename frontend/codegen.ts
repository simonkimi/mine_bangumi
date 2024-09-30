import path from "path";
import type { CodegenConfig } from "@graphql-codegen/cli";

const config: CodegenConfig = {
  overwrite: true,
  schema: [path.resolve(__dirname, "../graph/schema/**/*.graphql")],
  documents: "src/graph/**/*.graphql",
  generates: {
    "src/generate/": {
      preset: "client",
      plugins: [],
    },
    "./graphql.schema.json": {
      plugins: ["introspection"],
    },
  },
};

export default config;
