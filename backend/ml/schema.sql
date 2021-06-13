DROP TABLE IF EXISTS children_texts;

CREATE TABLE children_texts (
	children_id VARCHAR(36) NOT NULL,
	text TEXT NOT NULL,
	time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	INDEX (children_id)
);
