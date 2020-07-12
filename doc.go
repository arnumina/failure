/*
#######
##         ____     _ __
##        / __/__ _(_) /_ _________
##       / __/ _ `/ / / // / __/ -_)
##      /_/  \_,_/_/_/\_,_/_/  \__/
##
####### (c) 2020 Institut National de l'Audiovisuel ######################################## Archivage Numérique #######
*/

// Package failure allows to add context, debugging information to errors.
//
// So instead of the traditional:
//
//     if err != nil {
//         return err
//     }
//
// it can be interesting to generate a new error including the initial error
// while giving it some context:
//
//     if err := task.Execute(); err != nil {
//         return failure.New(err).Set("task", task.ID()).Msg("impossible to perform this task")
//     }
//
// It is of course possible to generate completely new errors :
//
//     if got != want {
//         return failure.New(nil).
//             Set("got", got).
//             Set("want", want).
//             Msg("the value received is not the right one")
//     }
//
package failure

/*
######################################################################################################## @(°_°)@ #######
*/
