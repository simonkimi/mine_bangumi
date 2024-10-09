/* eslint-disable */
import type { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
export type MakeEmpty<T extends { [key: string]: unknown }, K extends keyof T> = { [_ in K]?: never };
export type Incremental<T> = T | { [P in keyof T]?: P extends ' $fragmentName' | '__typename' ? T[P] : never };
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: { input: string; output: string; }
  String: { input: string; output: string; }
  Boolean: { input: boolean; output: boolean; }
  Int: { input: number; output: number; }
  Float: { input: number; output: number; }
};

export enum ApiStatusEnum {
  BadRequest = 'BAD_REQUEST',
  Cancel = 'CANCEL',
  DatabaseMigrationError = 'DATABASE_MIGRATION_ERROR',
  Forbidden = 'FORBIDDEN',
  FormValidationError = 'FORM_VALIDATION_ERROR',
  InternalServerError = 'INTERNAL_SERVER_ERROR',
  NotFound = 'NOT_FOUND',
  Success = 'SUCCESS',
  ThirdPartyApiError = 'THIRD_PARTY_API_ERROR',
  Timeout = 'TIMEOUT',
  Unauthorized = 'UNAUTHORIZED',
  UserCredentialsError = 'USER_CREDENTIALS_ERROR'
}

export type ConfigResult = {
  __typename?: 'ConfigResult';
  user: UserConfigResult;
};

export type Mutation = {
  __typename?: 'Mutation';
  /** 配置用户 */
  configUser: ConfigResult;
  /** 刷新API令牌 */
  refreshApiToken: UserConfigResult;
};


export type MutationConfigUserArgs = {
  input: UserConfigInput;
};

export type ParseAcgSourceInput = {
  parser: SourceParserEnum;
  source: Scalars['String']['input'];
};

export type ParseAcgSourceResult = {
  __typename?: 'ParseAcgSourceResult';
  files: Array<Scalars['String']['output']>;
  season: Scalars['Int']['output'];
  title: Scalars['String']['output'];
};

export type Query = {
  __typename?: 'Query';
  /** 刮削Acg数据源 */
  scraperDb: Array<ScrapeAcgResult>;
  /** 解析Acg数据源 */
  scraperSource: ParseAcgSourceResult;
};


export type QueryScraperDbArgs = {
  input: ScrapeAcgSourceInput;
};


export type QueryScraperSourceArgs = {
  input: ParseAcgSourceInput;
};

export type ScrapeAcgResult = {
  __typename?: 'ScrapeAcgResult';
  backdrop: Scalars['String']['output'];
  firstAirDate: Scalars['String']['output'];
  originalTitle: Scalars['String']['output'];
  overview: Scalars['String']['output'];
  poster: Scalars['String']['output'];
  scraper: ScraperEnum;
  seasons: Array<ScrapeAcgSeasonResult>;
  title: Scalars['String']['output'];
};

export type ScrapeAcgSeasonResult = {
  __typename?: 'ScrapeAcgSeasonResult';
  overview: Scalars['String']['output'];
  poster: Scalars['String']['output'];
  seasonId: Scalars['Int']['output'];
  title: Scalars['String']['output'];
};

export type ScrapeAcgSourceInput = {
  language: ScraperLanguage;
  scraper: ScraperEnum;
  title: Scalars['String']['input'];
};

export enum ScraperEnum {
  Tmdb = 'TMDB'
}

export enum ScraperLanguage {
  En = 'EN',
  Ja = 'JA',
  ZhHans = 'Zh_HANS',
  ZhHant = 'Zh_HANT'
}

export enum SourceParserEnum {
  Bangumi = 'BANGUMI'
}

export type UserConfigInput = {
  password?: InputMaybe<Scalars['String']['input']>;
  username?: InputMaybe<Scalars['String']['input']>;
};

export type UserConfigResult = {
  __typename?: 'UserConfigResult';
  token: Scalars['String']['output'];
  username: Scalars['String']['output'];
};

export type UserConfigDataFragment = { __typename?: 'UserConfigResult', username: string, token: string } & { ' $fragmentName'?: 'UserConfigDataFragment' };

export type ConfigDataFragment = { __typename?: 'ConfigResult', user: (
    { __typename?: 'UserConfigResult' }
    & { ' $fragmentRefs'?: { 'UserConfigDataFragment': UserConfigDataFragment } }
  ) } & { ' $fragmentName'?: 'ConfigDataFragment' };

export type ConfigUserMutationVariables = Exact<{
  input: UserConfigInput;
}>;


export type ConfigUserMutation = { __typename?: 'Mutation', configUser: (
    { __typename?: 'ConfigResult' }
    & { ' $fragmentRefs'?: { 'ConfigDataFragment': ConfigDataFragment } }
  ) };

export type RefreshApiTokenMutationVariables = Exact<{ [key: string]: never; }>;


export type RefreshApiTokenMutation = { __typename?: 'Mutation', refreshApiToken: (
    { __typename?: 'UserConfigResult' }
    & { ' $fragmentRefs'?: { 'UserConfigDataFragment': UserConfigDataFragment } }
  ) };

export const UserConfigDataFragmentDoc = {"kind":"Document","definitions":[{"kind":"FragmentDefinition","name":{"kind":"Name","value":"UserConfigData"},"typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"UserConfigResult"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"username"}},{"kind":"Field","name":{"kind":"Name","value":"token"}}]}}]} as unknown as DocumentNode<UserConfigDataFragment, unknown>;
export const ConfigDataFragmentDoc = {"kind":"Document","definitions":[{"kind":"FragmentDefinition","name":{"kind":"Name","value":"ConfigData"},"typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"ConfigResult"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"user"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"FragmentSpread","name":{"kind":"Name","value":"UserConfigData"}}]}}]}},{"kind":"FragmentDefinition","name":{"kind":"Name","value":"UserConfigData"},"typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"UserConfigResult"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"username"}},{"kind":"Field","name":{"kind":"Name","value":"token"}}]}}]} as unknown as DocumentNode<ConfigDataFragment, unknown>;
export const ConfigUserDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"ConfigUser"},"variableDefinitions":[{"kind":"VariableDefinition","variable":{"kind":"Variable","name":{"kind":"Name","value":"input"}},"type":{"kind":"NonNullType","type":{"kind":"NamedType","name":{"kind":"Name","value":"UserConfigInput"}}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"configUser"},"arguments":[{"kind":"Argument","name":{"kind":"Name","value":"input"},"value":{"kind":"Variable","name":{"kind":"Name","value":"input"}}}],"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"FragmentSpread","name":{"kind":"Name","value":"ConfigData"}}]}}]}},{"kind":"FragmentDefinition","name":{"kind":"Name","value":"UserConfigData"},"typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"UserConfigResult"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"username"}},{"kind":"Field","name":{"kind":"Name","value":"token"}}]}},{"kind":"FragmentDefinition","name":{"kind":"Name","value":"ConfigData"},"typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"ConfigResult"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"user"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"FragmentSpread","name":{"kind":"Name","value":"UserConfigData"}}]}}]}}]} as unknown as DocumentNode<ConfigUserMutation, ConfigUserMutationVariables>;
export const RefreshApiTokenDocument = {"kind":"Document","definitions":[{"kind":"OperationDefinition","operation":"mutation","name":{"kind":"Name","value":"RefreshApiToken"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"refreshApiToken"},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"FragmentSpread","name":{"kind":"Name","value":"UserConfigData"}}]}}]}},{"kind":"FragmentDefinition","name":{"kind":"Name","value":"UserConfigData"},"typeCondition":{"kind":"NamedType","name":{"kind":"Name","value":"UserConfigResult"}},"selectionSet":{"kind":"SelectionSet","selections":[{"kind":"Field","name":{"kind":"Name","value":"username"}},{"kind":"Field","name":{"kind":"Name","value":"token"}}]}}]} as unknown as DocumentNode<RefreshApiTokenMutation, RefreshApiTokenMutationVariables>;