/* eslint-disable */
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
