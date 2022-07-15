build:
	@echo Compiling GoMod...
	@go build --trimpath --buildmode=c-shared -ldflags "-extldflags -static" -o GoMod.dll .
	@move GoMod.dll "C:\Program Files (x86)\Steam\steamapps\common\VRChat\GoMods\"
	@echo Compiled GoMod.
