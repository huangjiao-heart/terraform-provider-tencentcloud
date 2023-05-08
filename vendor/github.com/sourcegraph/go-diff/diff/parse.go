	"path/filepath"
	"strconv"
// ParseMultiFileDiff parses a multi-file unified diff. It returns an error if
// parsing failed as a whole, but does its best to parse as many files in the
// case of per-file errors. If it cannot detect when the diff of the next file
// begins, the hunks are added to the FileDiff of the previous file.
	return &MultiFileDiffReader{reader: newLineReader(r)}
	reader *lineReader
	fd, _, err := r.ReadFileWithTrailingContent()
	return fd, err
}

// ReadFileWithTrailingContent reads the next file unified diff (including
// headers and all hunks) from r, also returning any trailing content. If there
// are no more files in the diff, it returns error io.EOF.
func (r *MultiFileDiffReader) ReadFileWithTrailingContent() (*FileDiff, string, error) {
				// Any non-diff content preceding a valid diff is included in the
				// extended headers of the following diff. In this way, mixed diff /
				// non-diff content can be parsed. Trailing non-diff content is
				// different: it doesn't make sense to return a FileDiff with only
				// extended headers populated. Instead, we return any trailing content
				// in case the caller needs it.
				trailing := ""
				if fd != nil {
					trailing = strings.Join(fd.Extended, "\n")
				}
				return nil, trailing, io.EOF
			return nil, "", err
			return fd, "", nil
			return nil, "", err
	// FileDiff is added/deleted file
	// No further collection of hunks needed
	if fd.NewName == "" {
		return fd, "", nil
	}

	line, err := r.reader.readLine()
	if err != nil && err != io.EOF {
		return fd, "", err
					return fd, "", nil
			return nil, "", err
	return fd, "", nil
	return &FileDiffReader{reader: &lineReader{reader: bufio.NewReader(r)}}
	reader *lineReader
		fd.OrigTime = origTime
		fd.NewTime = newTime
// timestamps). Or which starts with "Only in " with dir path and filename.
// "Only in" message is supported in POSIX locale: https://pubs.opengroup.org/onlinepubs/9699919799/utilities/diff.html#tag_20_34_10
	if r.fileHeaderLine != nil {
		if isOnlyMessage, source, filename := parseOnlyInMessage(r.fileHeaderLine); isOnlyMessage {
			return filepath.Join(string(source), string(filename)),
				"", nil, nil, nil
		}
	}
	unquotedOrigName, err := strconv.Unquote(origName)
	if err == nil {
		origName = unquotedOrigName
	}
	unquotedNewName, err := strconv.Unquote(newName)
	if err == nil {
		newName = unquotedNewName
	}

		line, err = r.reader.readLine()
		var ts time.Time
		ts, err = time.Parse(diffTimeParseLayout, parts[1])
			var err1 error
			ts, err1 = time.Parse(diffTimeParseWithoutTZLayout, parts[1])
			if err1 != nil {
				return "", nil, err
			}
			err = nil
	return fmt.Sprintf("overflowed into next file: %s", string(e))
			line, err = r.reader.readLine()
		// Reached message that file is added/deleted
		if isOnlyInMessage, _, _ := parseOnlyInMessage(line); isOnlyInMessage {
			r.fileHeaderLine = line // pass to readOneFileHeader (see fileHeaderLine field doc)
			return xheaders, nil
		}

// readQuotedFilename extracts a quoted filename from the beginning of a string,
// returning the unquoted filename and any remaining text after the filename.
func readQuotedFilename(text string) (value string, remainder string, err error) {
	if text == "" || text[0] != '"' {
		return "", "", fmt.Errorf(`string must start with a '"': %s`, text)
	}

	// The end quote is the first quote NOT preceeded by an uneven number of backslashes.
	numberOfBackslashes := 0
	for i, c := range text {
		if c == '"' && i > 0 && numberOfBackslashes%2 == 0 {
			value, err = strconv.Unquote(text[:i+1])
			remainder = text[i+1:]
			return
		} else if c == '\\' {
			numberOfBackslashes++
		} else {
			numberOfBackslashes = 0
		}
	}
	return "", "", fmt.Errorf(`end of string found while searching for '"': %s`, text)
}

// parseDiffGitArgs extracts the two filenames from a 'diff --git' line.
// Returns false on syntax error, true if syntax is valid. Even with a
// valid syntax, it may be impossible to extract filenames; if so, the
// function returns ("", "", true).
func parseDiffGitArgs(diffArgs string) (string, string, bool) {
	length := len(diffArgs)
	if length < 3 {
		return "", "", false
	}

	if diffArgs[0] != '"' && diffArgs[length-1] != '"' {
		// Both filenames are unquoted.
		firstSpace := strings.IndexByte(diffArgs, ' ')
		if firstSpace <= 0 || firstSpace == length-1 {
			return "", "", false
		}

		secondSpace := strings.IndexByte(diffArgs[firstSpace+1:], ' ')
		if secondSpace == -1 {
			if diffArgs[firstSpace+1] == '"' {
				// The second filename begins with '"', but doesn't end with one.
				return "", "", false
			}
			return diffArgs[:firstSpace], diffArgs[firstSpace+1:], true
		}

		// One or both filenames contain a space, but the names are
		// unquoted. Here, the 'diff --git' syntax is ambiguous, and
		// we have to obtain the filenames elsewhere (e.g. from the
		// hunk headers or extended headers). HOWEVER, if the file
		// is newly created and empty, there IS no other place to
		// find the filename. In this case, the two filenames are
		// identical (except for the leading 'a/' prefix), and we have
		// to handle that case here.
		first := diffArgs[:length/2]
		second := diffArgs[length/2+1:]

		// If the two strings could be equal, based on length, proceed.
		if length%2 == 1 {
			// If the name minus the a/ b/ prefixes is equal, proceed.
			if len(first) >= 3 && first[1] == '/' && first[1:] == second[1:] {
				return first, second, true
			}
			// If the names don't have the a/ and b/ prefixes and they're equal, proceed.
			if !(first[:2] == "a/" && second[:2] == "b/") && first == second {
				return first, second, true
			}
		}

		// The syntax is (unfortunately) valid, but we could not extract
		// the filenames.
		return "", "", true
	}

	if diffArgs[0] == '"' {
		first, remainder, err := readQuotedFilename(diffArgs)
		if err != nil || len(remainder) < 2 || remainder[0] != ' ' {
			return "", "", false
		}
		if remainder[1] == '"' {
			second, remainder, err := readQuotedFilename(remainder[1:])
			if remainder != "" || err != nil {
				return "", "", false
			}
			return first, second, true
		}
		return first, remainder[1:], true
	}

	// In this case, second argument MUST be quoted (or it's a syntax error)
	i := strings.IndexByte(diffArgs, '"')
	if i == -1 || i+2 >= length || diffArgs[i-1] != ' ' {
		return "", "", false
	}

	second, remainder, err := readQuotedFilename(diffArgs[i:])
	if remainder != "" || err != nil {
		return "", "", false
	}
	return diffArgs[:i-1], second, true
}

	lineCount := len(fd.Extended)
	if lineCount > 0 && !strings.HasPrefix(fd.Extended[0], "diff --git ") {
		return false
	}
	lineHasPrefix := func(idx int, prefix string) bool {
		return strings.HasPrefix(fd.Extended[idx], prefix)
	}

	linesHavePrefixes := func(idx1 int, prefix1 string, idx2 int, prefix2 string) bool {
		return lineHasPrefix(idx1, prefix1) && lineHasPrefix(idx2, prefix2)
	}

	isCopy := (lineCount == 4 && linesHavePrefixes(2, "copy from ", 3, "copy to ")) ||
		(lineCount == 6 && linesHavePrefixes(2, "copy from ", 3, "copy to ") && lineHasPrefix(5, "Binary files ")) ||
		(lineCount == 6 && linesHavePrefixes(1, "old mode ", 2, "new mode ") && linesHavePrefixes(4, "copy from ", 5, "copy to "))

	isRename := (lineCount == 4 && linesHavePrefixes(2, "rename from ", 3, "rename to ")) ||
		(lineCount == 5 && linesHavePrefixes(2, "rename from ", 3, "rename to ") && lineHasPrefix(4, "Binary files ")) ||
		(lineCount == 6 && linesHavePrefixes(2, "rename from ", 3, "rename to ") && lineHasPrefix(5, "Binary files ")) ||
		(lineCount == 6 && linesHavePrefixes(1, "old mode ", 2, "new mode ") && linesHavePrefixes(4, "rename from ", 5, "rename to "))

	isDeletedFile := (lineCount == 3 || lineCount == 4 && lineHasPrefix(3, "Binary files ") || lineCount > 4 && lineHasPrefix(3, "GIT binary patch")) &&
		lineHasPrefix(1, "deleted file mode ")

	isNewFile := (lineCount == 3 || lineCount == 4 && lineHasPrefix(3, "Binary files ") || lineCount > 4 && lineHasPrefix(3, "GIT binary patch")) &&
		lineHasPrefix(1, "new file mode ")

	isModeChange := lineCount == 3 && linesHavePrefixes(1, "old mode ", 2, "new mode ")

	isBinaryPatch := lineCount == 3 && lineHasPrefix(2, "Binary files ") || lineCount > 3 && lineHasPrefix(2, "GIT binary patch")

	if !isModeChange && !isCopy && !isRename && !isBinaryPatch && !isNewFile && !isDeletedFile {
		return false
	}

	var success bool
	fd.OrigName, fd.NewName, success = parseDiffGitArgs(fd.Extended[0][len("diff --git "):])
	if isNewFile {
	}
	if isDeletedFile {

	// For ambiguous 'diff --git' lines, try to reconstruct filenames using extended headers.
	if success && (isCopy || isRename) && fd.OrigName == "" && fd.NewName == "" {
		diffArgs := fd.Extended[0][len("diff --git "):]

		tryReconstruct := func(header string, prefix string, whichFile int, result *string) {
			if !strings.HasPrefix(header, prefix) {
				return
			}
			rawFilename := header[len(prefix):]

			// extract the filename prefix (e.g. "a/") from the 'diff --git' line.
			var prefixLetterIndex int
			if whichFile == 1 {
				prefixLetterIndex = 0
			} else if whichFile == 2 {
				prefixLetterIndex = len(diffArgs) - len(rawFilename) - 2
			}
			if prefixLetterIndex < 0 || diffArgs[prefixLetterIndex+1] != '/' {
				return
			}

			*result = diffArgs[prefixLetterIndex:prefixLetterIndex+2] + rawFilename
		}

		for _, header := range fd.Extended {
			tryReconstruct(header, "copy from ", 1, &fd.OrigName)
			tryReconstruct(header, "copy to ", 2, &fd.NewName)
			tryReconstruct(header, "rename from ", 1, &fd.OrigName)
			tryReconstruct(header, "rename to ", 2, &fd.NewName)
		}
	}
	return success

	// ErrBadOnlyInMessage is when a file have a malformed `only in` message
	// Should be in format `Only in {source}: {filename}`
	ErrBadOnlyInMessage = errors.New("bad 'only in' message")
	return &HunksReader{reader: &lineReader{reader: bufio.NewReader(r)}}
	reader *lineReader
			line, err = r.reader.readLine()

			// If the line starts with `---` and the next one with `+++` we're
			// looking at a non-extended file header and need to abort.
			if bytes.HasPrefix(line, []byte("---")) {
				ok, err := r.reader.nextLineStartsWith("+++")
				if err != nil {
					return r.hunk, err
				}
				if ok {
					ok2, _ := r.reader.nextNextLineStartsWith(string(hunkPrefix))
					if ok2 {
						return r.hunk, &ParseError{r.line, r.offset, &ErrBadHunkLine{Line: line}}
					}
				}
			}

			// If the line starts with the hunk prefix, this hunk is complete.
				// But we've already read in the next hunk's
// parseOnlyInMessage checks if line is a "Only in {source}: {filename}" and returns source and filename
func parseOnlyInMessage(line []byte) (bool, []byte, []byte) {
	if !bytes.HasPrefix(line, onlyInMessagePrefix) {
		return false, nil, nil
	}
	line = line[len(onlyInMessagePrefix):]
	idx := bytes.Index(line, []byte(": "))
	if idx < 0 {
		return false, nil, nil
	}
	return true, line[:idx], line[idx+2:]
}
