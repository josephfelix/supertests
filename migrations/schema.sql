--
-- PostgreSQL database dump
--

-- Dumped from database version 11.5 (Debian 11.5-1.pgdg90+1)
-- Dumped by pg_dump version 11.10 (Debian 11.10-0+deb10u1)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: results; Type: TABLE; Schema: public; Owner: supertests
--

CREATE TABLE public.results (
    id integer NOT NULL,
    userid uuid NOT NULL,
    guid uuid NOT NULL,
    result character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.results OWNER TO supertests;

--
-- Name: results_id_seq; Type: SEQUENCE; Schema: public; Owner: supertests
--

CREATE SEQUENCE public.results_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.results_id_seq OWNER TO supertests;

--
-- Name: results_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: supertests
--

ALTER SEQUENCE public.results_id_seq OWNED BY public.results.id;


--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: supertests
--

CREATE TABLE public.schema_migration (
    version character varying(14) NOT NULL
);


ALTER TABLE public.schema_migration OWNER TO supertests;

--
-- Name: tests; Type: TABLE; Schema: public; Owner: supertests
--

CREATE TABLE public.tests (
    id integer NOT NULL,
    title character varying(255) NOT NULL,
    cover character varying(255) NOT NULL,
    slug character varying(255) NOT NULL,
    description character varying(255) NOT NULL,
    message character varying(255) NOT NULL,
    plugin character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.tests OWNER TO supertests;

--
-- Name: tests_id_seq; Type: SEQUENCE; Schema: public; Owner: supertests
--

CREATE SEQUENCE public.tests_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.tests_id_seq OWNER TO supertests;

--
-- Name: tests_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: supertests
--

ALTER SEQUENCE public.tests_id_seq OWNED BY public.tests.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: supertests
--

CREATE TABLE public.users (
    id uuid NOT NULL,
    name character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    gender character varying(255),
    age integer,
    cover character varying(255),
    favorite_athletes character varying(255),
    favorite_teams character varying(255),
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE public.users OWNER TO supertests;

--
-- Name: results id; Type: DEFAULT; Schema: public; Owner: supertests
--

ALTER TABLE ONLY public.results ALTER COLUMN id SET DEFAULT nextval('public.results_id_seq'::regclass);


--
-- Name: tests id; Type: DEFAULT; Schema: public; Owner: supertests
--

ALTER TABLE ONLY public.tests ALTER COLUMN id SET DEFAULT nextval('public.tests_id_seq'::regclass);


--
-- Name: results results_pkey; Type: CONSTRAINT; Schema: public; Owner: supertests
--

ALTER TABLE ONLY public.results
    ADD CONSTRAINT results_pkey PRIMARY KEY (id);


--
-- Name: tests tests_pkey; Type: CONSTRAINT; Schema: public; Owner: supertests
--

ALTER TABLE ONLY public.tests
    ADD CONSTRAINT tests_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: supertests
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: supertests
--

CREATE UNIQUE INDEX schema_migration_version_idx ON public.schema_migration USING btree (version);


--
-- Name: users_email_idx; Type: INDEX; Schema: public; Owner: supertests
--

CREATE UNIQUE INDEX users_email_idx ON public.users USING btree (email);


--
-- PostgreSQL database dump complete
--

