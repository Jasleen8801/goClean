import { StyleSheet, View } from 'react-native';
import StackNavigator from './src/components/StackNavigator';
import { ModalPortal } from 'react-native-modals';

export default function App() {
  return (
    <>
      <View style={styles.container}>
        <StackNavigator />
        <ModalPortal />
      </View>
    </>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
  }
});
