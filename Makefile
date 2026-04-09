NAME = intersect

# ---------- Helper scripts ----------

ORG2NW    := bash scripts/org2nw.sh
PRETANGLE := awk -f scripts/preTangle.awk

# ---------- Basic tangling ----------

all: $(NAME).go lang_actions

$(NAME).go: $(NAME).org
	$(ORG2NW) $(NAME).org | $(PRETANGLE) | notangle -R$(NAME).go > $(NAME).go

# ---------- Basic make subcommands ----------

.PHONY: doc clean

doc:
	make -C doc

clean:
	rm -f $(NAME) *.go
	make clean -C doc

publish:
	if mountpoint -q ~/owncloud; then \
		cp doc/$(NAME)Doc.pdf ~/owncloud/github_docs; \
	fi

# ---------- Language actions area ----------

lang_actions: $(NAME).go go.mod go.sum
	gofmt -w $(NAME).go
	go build $(NAME).go

go.mod:
	go mod init $(NAME).go

go.sum: go.mod $(NAME).go
	go mod tidy
