/// PlayerID identifies a Player by their position within a game.
enum PlayerID { noPlayer, player1, player2 }

String playerIDToString(PlayerID id) {
  switch (id) {
    case PlayerID.noPlayer:
      return 'no player';
    case PlayerID.player1:
      return 'player 1';
    case PlayerID.player2:
      return 'player 2';
  }
}

PlayerID stringToPlayerID(String x) {
  switch (x) {
    case 'no player':
      return PlayerID.noPlayer;
    case 'player 1':
      return PlayerID.player1;
    case 'player 2':
      return PlayerID.player2;
  }
}

/// Player of a game.
class Player {
  /// name of the Player.
  final String name;

  /// description of the Player.
  final String description;

  Map<String, dynamic> arguments = {};

  /// Player constructor initializes the Player's name and description.
  Player(this.name, {this.description: '', this.arguments}) {
    if (arguments == null) {
      arguments = {};
    }
  }
}
