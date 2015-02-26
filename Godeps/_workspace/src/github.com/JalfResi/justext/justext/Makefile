#include $(GOROOT)/src/Make.inc

#TARG=ourscienceistight/gojustext
#GOFILES=\
#	preprocess.go\
#	paragraphObjectModel.go\
#	htmlRenderer.go\
#	classifyParagraphs.go\
#	reviseParagraphClassification.go\
#	reader.go\
#	writer.go\
#	stoplists.go\
#	stoplists/*.go\
#	defaultTemplate.go\
#	detailedTemplate.go\
#
#include $(GOROOT)/src/Make.pkg

objects = stoplists/*.go

edit: $(objects)

stoplists/%.go:
	mkdir stoplists/$*
	cp stoplists/template stoplists/$*/main.go
	sed -i 's/<language>/$*/g' stoplists/$*/main.go
	wget -Ostoplists/$*/data.txt http://justext.googlecode.com/svn/trunk/justext/stoplists/$*.txt
	go-bindata -i=stoplists/$*/data.txt -o="stoplists/$*/data.go" -f="$*Stoplist" -p="justext/stoplists/$*"

defaultTemplate.go:
	go-bindata -i=default.template -o="../defaultTemplate.go" -f="DefaultTemplate" -p="justext"

detailedTemplate.go:
	go-bindata -i=detailed.template -o="../detailedTemplate.go" -f="DetailedTemplate" -p="justext"
