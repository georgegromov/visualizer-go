create table if not exists users (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  username VARCHAR(32) NOT NULL,
  password_hash VARCHAR(255) NOT NULL,
  role TEXT NOT NULL DEFAULT 'User' CHECK (role IN ('Admin', 'User', 'Customer')),
  updated_at TIMESTAMPTZ,
  created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

create table if not exists templates (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	name VARCHAR(64),
  description VARCHAR(255),
  is_deleted BOOLEAN DEFAULT false,
	creator_id UUID,
	
	updated_at TIMESTAMPTZ,
  created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	
	CONSTRAINT fk_template_creator FOREIGN KEY (creator_id) REFERENCES users(id)
);

create table if not exists canvases (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name TEXT,
  description TEXT,
  template_id UUID NOT NULL,
  updated_at TIMESTAMP,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  CONSTRAINT fk_canvas_template FOREIGN KEY (template_id) REFERENCES templates(id) ON DELETE CASCADE
);

create table if not exists charts (
  	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  	name TEXT,
  	type TEXT NOT NULL,
  	canvas_id UUID NOT NULL,
  	updated_at TIMESTAMP,
  	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
  	CONSTRAINT fk_chart_canvas FOREIGN KEY (canvas_id) REFERENCES canvases(id) ON DELETE CASCADE
);

create table if not exists measurements (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	content jsonb NOT NULL,
	chart_id UUID NOT NULL,
	updated_at TIMESTAMP,
  created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	CONSTRAINT fk_chart_measurement FOREIGN KEY (chart_id) REFERENCES charts(id) ON DELETE CASCADE
);

create table if not exists dashboards (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	name VARCHAR(64) NOT NULL,
  description VARCHAR(255),
	tenant TEXT NOT NULL DEFAULT 'Imby' CHECK (tenant IN ('Imby', 'EF'))
	is_published BOOL NOT NULL DEFAULT false,
	share_id UUID NOT NULL DEFAULT gen_random_uuid(),
	creator_id UUID NOT NULL,
	template_id UUID,
	
	view_count INTEGER DEFAULT 0,
	
  last_viewed_at TIMESTAMP,
	updated_at TIMESTAMPTZ,
  created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
	
	CONSTRAINT fk_dashboard_template FOREIGN KEY (template_id) REFERENCES templates(id),
  CONSTRAINT fk_dashboard_creator FOREIGN KEY (creator_id) REFERENCES users(id)
);

create table if not exists diagrams (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	name VARCHAR(64) NOT NULL,
	content json,
	updated_at TIMESTAMPTZ,
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);