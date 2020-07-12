/*
#######
##         ____     _ __
##        / __/__ _(_) /_ _________
##       / __/ _ `/ / / // / __/ -_)
##      /_/  \_,_/_/_/\_,_/_/  \__/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

package failure

import (
	"testing"
)

func TestUnexpected(t *testing.T) {
	want := _unexpectedErrMessage
	if got := Unexpected().Error(); got != want {
		t.Errorf("Unexpected(): got: %s, want %s", got, want)
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
